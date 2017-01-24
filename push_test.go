package jmapns

import (
	"testing"
)

var payload = Payload{
	Sound: "sound.aif",
}

var notification = Notification{
	DeviceToken: "deadbeef",
	Payload:     payload,
}

func TestNoCert(t *testing.T) {
	clearAPNsCertificate()

}
