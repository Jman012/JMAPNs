package main

import (
	"JMAPNs"
	"flag"
	"fmt"
	"time"
)

var certFile = flag.String("cert", "", "Path to APNs Certificate")
var keyFile = flag.String("key", "", "Path to APNs Key Certificate")
var token = flag.String("t", "", "Sample device token")

func main() {
	flag.Parse()

	if *certFile == "" || *keyFile == "" || *token == "" {
		fmt.Println("Bad arguments")
		return
	}

	JMAPNs.Development()
	JMAPNs.EnableSuccessResponses()
	err := JMAPNs.LoadAPNsCertificate(*certFile, *keyFile)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	payload := JMAPNs.NewPayload()
	payload.Alert.SetBody("This is a test notification").SetTitle("The title")

	notification := JMAPNs.Notification{
		Payload:     *payload,
		DeviceToken: *token,
	}

	go func() {
		for resp := range JMAPNs.ResponseChannel {
			fmt.Printf("Received response: %#v\n", resp)
		}
	}()

	JMAPNs.SendChannel <- &notification

	time.Sleep(5 * time.Second)
}
