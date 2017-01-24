package jmapns

import (
	"testing"
)

func TestEmptyPayload(t *testing.T) {
	p := NewPayload()
	s := p.String()

	expected := `{"aps":{}}`

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}

func TestSilent(t *testing.T) {
	p := NewPayload().SetSilent()
	s := p.String()

	expected := `{"aps":{"content-available":1}}`

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}

func TestAppSpecificPayload(t *testing.T) {
	p := NewPayload().SetSound("sound.aif")
	p.SetAppSpecific("acme2", "test")
	s := p.String()

	expected := `{"acme2":"test","aps":{"sound":"sound.aif"}}`

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}

func TestBadge(t *testing.T) {
	p := NewPayload().SetBadge(2)
	s := p.String()

	expected := `{"aps":{"badge":2}}`

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}

func TestArgs(t *testing.T) {
	p := NewPayload()
	p.Alert.SetTitleLocalizedArguments([]string{"arg1", "arg2"})
	s := p.String()

	expected := `{"aps":{"alert":{"title-loc-args":["arg1","arg2"]}}}`

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}

func TestPartialPayload(t *testing.T) {
	p := NewPayload().SetSound("test.aif").SetThreadId("1")
	p.Alert.SetBody("the body").SetLaunchImage("image.png")

	expected := `{"aps":{"alert":{"body":"the body","launch-image":"image.png"},"sound":"test.aif","thread-id":"1"}}`

	s := p.String()

	if s != expected {
		t.Errorf("Expected %#v, got %#v", expected, s)
	}
}
