package jmapns

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Payload struct {
	Alert              *Alert                 `json:"alert,omitempty"`
	Badge              json.Number            `json:"badge,omitempty,Number"` // We need 0 value
	Sound              string                 `json:"sound,omitempty"`
	SilentNotification numberedBool           `json:"content-available,omitempty"`
	Category           string                 `json:"category,omitempty"`
	ThreadId           string                 `json:"thread-id,omitempty"`
	AppSpecific        map[string]interface{} `json:"-"`
}

func NewPayload() *Payload {
	p := &Payload{
		Alert: &Alert{},
	}
	return p
}

func (p *Payload) String() string {
	return string(p.Bytes())
}

func (p *Payload) Bytes() []byte {
	// Prepare the payload for correct packaging first
	entirePayload := make(map[string]interface{})
	emptyAlert := Alert{}
	if reflect.DeepEqual(*p.Alert, emptyAlert) {
		p.Alert = nil
	}
	entirePayload["aps"] = p

	for key, val := range p.AppSpecific {
		entirePayload[key] = val
	}

	// Then convert it
	bytes, err := json.Marshal(entirePayload)

	if err != nil {
		// Unrecoverable error, but should never happen
		fmt.Printf("Unexpected error converting Payload to string: %v", err)
		panic(1)
	}
	return bytes
}

func (p *Payload) SetBadge(num int) *Payload {
	p.Badge = json.Number(strconv.Itoa(num))
	return p
}

func (p *Payload) SetSound(sound string) *Payload {
	p.Sound = sound
	return p
}

func (p *Payload) SetSilent() *Payload {
	p.SilentNotification = true
	return p
}

func (p *Payload) SetCategory(cat string) *Payload {
	p.Category = cat
	return p
}

func (p *Payload) SetThreadId(id string) *Payload {
	p.ThreadId = id
	return p
}

func (p *Payload) SetAppSpecific(key, val string) *Payload {
	if p.AppSpecific == nil {
		p.AppSpecific = make(map[string]interface{})
	}
	p.AppSpecific[key] = val
	return p
}

type Alert struct {
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

func (a *Alert) SetTitle(title string) *Alert {
	a.Title = title
	return a
}

func (a *Alert) SetBody(body string) *Alert {
	a.Body = body
	return a
}

func (a *Alert) SetTitleLocalizedKey(key string) *Alert {
	a.TitleLocKey = key
	return a
}

func (a *Alert) SetTitleLocalizedArguments(args []string) *Alert {
	a.TitleLocArgs = args
	return a
}

func (a *Alert) SetActionLocalizedKey(key string) *Alert {
	a.ActionLocKey = key
	return a
}

func (a *Alert) SetLocalizedKey(key string) *Alert {
	a.LocKey = key
	return a
}

func (a *Alert) SetLocalizedArguments(args []string) *Alert {
	a.LocArgs = args
	return a
}

func (a *Alert) SetLaunchImage(image string) *Alert {
	a.LaunchImage = image
	return a
}

// For content-available
// Value should only ever be 1, so to enforce the type is bool
// And its MarshalJSON will put the number 1 in the JSON
type numberedBool bool

func (nb numberedBool) MarshalJSON() ([]byte, error) {
	if nb == true {
		return json.Marshal(1)
	} else {
		return json.Marshal(0)
	}
}
