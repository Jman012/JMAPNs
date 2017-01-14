package JMAPNs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Notification struct {
	Payload     Payload
	DeviceToken string
	ID          string
	Expiration  *time.Time
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
	DeviceToken  string
	ID           string
	Status       ResponseStatus
	ResponseBody string
	LocalError   error
}

var SendChannel = make(chan *Notification)
var ResponseChannel = make(chan Response)
var SuccessChannel = make(chan Response)

var successResponse = false

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
		return fmt.Errorf("nil Notification")
	}
	if http2Transport == nil || currentClient == nil {
		return fmt.Errorf("APNs certificate not loaded")
	}

	// Convert the payload into JSON, handling errors
	payloadBytes := not.Payload.Bytes()
	if len(payloadBytes) > MaximumPayloadSize {
		return fmt.Errorf("payload too large, expected %v was %v", MaximumPayloadSize, len(payloadBytes))
	}

	// Make the request
	url := fmt.Sprintf("%s/3/device/%s", apnsEndPoint, not.DeviceToken)
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return fmt.Errorf("unexpected error creating HTTP/2 request: %v", err)
	}
	not.applyHeaders(req)

	// Perform the request
	resp, err := currentClient.Do(req)
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

func EnableSuccessResponses() {
	successResponse = true
}

func DisableSuccessResponses() {
	successResponse = false
}

func parseResponse(resp *http.Response, not *Notification) error {
	defer resp.Body.Close()

	if successResponse == false && ResponseStatus(resp.StatusCode) == RespSuccess {
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	apnsResp := Response{
		DeviceToken:  not.DeviceToken,
		ID:           resp.Header.Get("apns-id"),
		Status:       ResponseStatus(resp.StatusCode),
		ResponseBody: string(body),
	}

	ResponseChannel <- apnsResp
	return nil
}

func (n *Notification) applyHeaders(req *http.Request) {
	if n.ID != "" {
		req.Header.Set("apns-id", n.ID)
	}

	if n.Expiration != nil {
		req.Header.Set("apns-expiration", strconv.FormatInt(n.Expiration.Unix(), 10))
	}

	if n.Priority == 10 || n.Priority == 5 {
		req.Header.Set("apns-priority", strconv.Itoa(n.Priority))
	}

	if n.Topic != "" {
		req.Header.Set("apns-topic", n.Topic)
	}

	if n.CollapseID != "" {
		req.Header.Set("apns-collapse-id", n.CollapseID)
	}
}
