package main

import (
	"flag"
	"fmt"
	"jmapns"
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

	jmapns.Development()
	jmapns.EnableSuccessResponses()
	err := jmapns.LoadAPNsCertificate(*certFile, *keyFile)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	payload := jmapns.NewPayload()
	payload.Alert.SetBody("This is a test notification").SetTitle("The title")

	notification := jmapns.Notification{
		Payload:     *payload,
		DeviceToken: jmapns.Token(*token),
	}

	go func() {
		for resp := range jmapns.ResponseChannel {
			fmt.Printf("Received response: %#v\n", resp)
		}
	}()

	go func() {
		for resp := range jmapns.SuccessChannel {
			fmt.Printf("Successful push: %#v\n", resp)
		}
	}()

	jmapns.SendChannel <- &notification

	time.Sleep(5 * time.Second)
}
