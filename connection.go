package JMAPNs

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	"os"
)

const (
	ProductionEndPoint  = "https://api.push.apple.com"
	DevelopmentEndPoint = "https://api.development.push.apple.com"
)

type APNsClient http.Client

func (a APNsClient) Do(req *http.Request) (*http.Response, error) {
	c := http.Client(a)
	return c.Do(req)
}

var apnsEndPoint string = ProductionEndPoint
var currentClient *APNsClient

func Production() {
	// TODO: When set, reconnect
	apnsEndPoint = ProductionEndPoint
}

func Development() {
	// TODO: When set, reconnect
	apnsEndPoint = DevelopmentEndPoint
}

var http2Transport *http.Transport = nil

func MustLoadAPNsCertificate(certFilePath, keyFilePath string) {
	if err := LoadAPNsCertificate(certFilePath, keyFilePath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func LoadAPNsCertificate(certFilePath, keyFilePath string) error {

	// Load the actual certificate files
	certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return fmt.Errorf("error loading certifcate and key files: %v", err)
	}

	// Then make a config and transport for future http.Clients to use
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	err = http2.ConfigureTransport(transport)
	if err != nil {
		return fmt.Errorf("error configuring HTTP/2 client, are you using Go >1.6?: %v", err)
	}

	// Set our global with the new, re-usable transport
	http2Transport = transport
	err = newAPNsClient()
	if err != nil {
		return fmt.Errorf("unexpected error creating HTTP/2 client")
	}
	return nil
}

func clearAPNsCertificate() {
	http2Transport = nil
}

func newAPNsClient() error {

	if http2Transport == nil {
		return fmt.Errorf("error: could not create APNs client, you did not load the certificate")
	}

	currentClient = &APNsClient{Transport: http2Transport}
	return nil
}
