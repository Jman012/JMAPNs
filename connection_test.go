package jmapns

import (
	"io/ioutil"
	"testing"
)

const (
	certFile  = "ExtraOrdinaryPush-Cert.pem"
	keyFile   = "ExtraOrdinaryPush-Key-NoEncryption.pem"
	wrongFile = "blank.txt"
)

func TestBadCert(t *testing.T) {
	// No files given
	clearAPNsCertificate()
	err := LoadAPNsCertificate("", "")
	if err == nil {
		t.Error("Didn't get error with no cert")
	}

	// Actual file, but not cert contents
	clearAPNsCertificate()
	err = LoadAPNsCertificate(wrongFile, wrongFile)
	if err == nil {
		t.Error("Didn't get error with invalid cert")
	}
}

func TestGoodCert(t *testing.T) {
	clearAPNsCertificate()
	err := LoadAPNsCertificate(certFile, keyFile)
	if err != nil {
		t.Errorf("Unexpected error loading valid cert: %v", err)
	}
}

func TestClientWorks(t *testing.T) {
	Development()
	clearAPNsCertificate()
	err := LoadAPNsCertificate(certFile, keyFile)
	if err != nil {
		t.Errorf("Unexpected setup error loading cert: %v", err)
	}

	resp, err := currentClient.Get(apnsEndPoint + "/3/device/invaliddevice")
	if err != nil {
		t.Errorf("Unexpected error communicating with APNs: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Logf("Resonse header: %#v\n", resp)
	t.Logf("Response body: %v\n", string(body))
}
