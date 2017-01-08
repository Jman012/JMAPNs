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

var apnsEndPoint string = ProductionEndPoint

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
	return nil
}

func clearAPNsCertificate() {
	http2Transport = nil
}

func NewAPNsClient() (*http.Client, error) {

	if http2Transport == nil {
		return nil, fmt.Errorf("error: could not create APNs client, you did not load the certificate")
	}

	client := &http.Client{Transport: http2Transport}

	return client, nil
}
