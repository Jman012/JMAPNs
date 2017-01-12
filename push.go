package JMAPNs

import (
	"fmt"
	"strconv"
	"time"
)

type Notification struct {
	Payload     Payload
	DeviceToken string
	ID          string
	Expiration  time.Time
	Priority    int
	Topic       string
	CollapseID  string
}

/*
	A Response object is sent on the response channel only in the case of some
	kind of error. If the error originated in constructing the HTTP request,
	then LocalError will have an error value set.

	However, if the request was sent and we got a response from APNs and that
	response had an error (not 200 status) then the other appropriate fields
	will have the error information.

	Successful notification receipts are not returned on the response channel.
	The device token will always be included, in order for the caller to know
	for what device an error occurred.
*/
type Response struct {
	DeviceToken string
	ID          string
	Status      ResponseStatus
	LocalError  error
}

var SendChannel = make(chan *Notification)
var ResponseChannel = make(chan Response)

func init() {
	go sender()
}

func sender() {
	for not := range SendChannel {
		if err := push(not); err != nil {
			ResponseChannel <- Response{DeviceToken: not.DeviceToken, LocalError: err}
		}
	}
}

func push(not *Notification) error {

	// First, sanity checks
	if not == nil {
		return fmt.Error("nil Notification")
	}
	if http2Transport == nil || apnsClient == nil {
		return fmt.Error("APNs certificate not loaded")
	}

	// Construct and send the notification and payload
	url := fmt.Sprintf("%s/3/device/%s", apnsEndPoint, not.DeviceToken)
	notBytes := not.Bytes()
	if len(notBytes) > MaximumPayloadSize {
		return fmt.Errorf("payload too large, expected %v was %v", MaximumPayloadSize, len(notBytes))
	}
	req := http.NewRequest("POST", url, bytes.NewReader(notBytes))
	not.applyHeaders(req)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP/2 request: %v", err)
	}

	// Handle the server's response
	err = parseResponse(resp, not)
	if err != nil {
		return err
	}

	return nil
}

func parseResponse(resp *http.Response, not *Notification) Response {
	if ResponseStatus(resp.StatusCode) != Success {
		apnsResp := &Response{
			DeviceToken: not.DeviceToken,
			ID:          resp.Header.Get("apns-id"),
			Status:      ResponseStatus(resp.StatusCode),
		}
	}
}

func (n *Notification) applyHeaders(req *http.Request) {
	if n.ID != "" {
		req.Set("apns-id", n.ID)
	}

	if n.Expiration != nil {
		req.Set("apns-expiration", strconv.Itoa(n.Expiration.Unix))
	}

	if n.Priority != 10 && n.Priority != 5 {
		req.Set("apns-priority", strconv.Itoa(n.Priority))
	}

	if n.Topic != "" {
		req.Set("apns-topic", n.Topic)
	}

	if n.CollapseId != "" {
		req.Set("apns-collapse-id", n.CollapseId)
	}
}
