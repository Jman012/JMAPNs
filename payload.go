package JMAPNs

import (
	"encoding/json"
)

const (
	MaximumSize = 4 * 1024
)

type Payload struct {
	Alert              Alert                  `json:"alert,omitempty"`
	Badge              json.Number            `json:"badge,omitempty,Number"` // We need 0 value
	Sound              string                 `json:"sound,omitempty"`
	SilentNotification numberedBool           `json:"content-available,omitempty"`
	Category           string                 `json:"category,omitempty"`
	ThreadId           string                 `json:"thread-id,omitempty"`
	AppSpecific        map[string]interface{} `json:"-"`
}

func NewPayload() *Payload {
	p = &Payload{AppSpecific: make(map[string]interface{})}
	return p
}

func (p *Payload) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		// Unrecoverable error, but should never happen
		panic(1)
	}
	return string(bytes)
}

func (p *Payload) Badge(num int) *Payload {
	p.Badge = json.Number(num)
	return p
}

func (p *Payload) Sound(sound string) *Payload {
	p.Sound = sound
	return p
}

func (p *Payload) RemoteNotification() *Payload {
	p.SilentNotification = true
	return p
}

func (p *Payload) Category(cat string) *Payload {
	p.Category = cat
	return p
}

func (p *Payload) ThreadId(id string) *Payload {
	p.ThreadId = id
	return p
}

func (p *Payload) MarshalJSON() ([]byte, error) {
	entirePayload := make(map[string]interface{})
	entirePayload["aps"] = p

	for key, val := range p.AppSpecific {
		entirePayload[key] = val
	}

	return json.Marshal(entirePayload)
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

func (a *Alert) Title(title string) *Alert {
	a.Title = title
	return a
}

func (a *Alert) Body(body string) *Alert {
	a.Body = body
	return a
}

func (a *Alert) TitleLocalizedKey(key string) *Alert {
	a.TitleLocKey = key
	return a
}

func (a *Alert) TitleLocalizedArguments(args []string) *Alert {
	a.TitleLocArgs = args
	return a
}

func (a *Alert) ActionLocalizedKey(key string) *Alert {
	a.ActionLocKey = key
	return a
}

func (a *Alert) LocalizedKey(key string) *Alert {
	a.LocKey = key
	return a
}

func (a *Alert) LocalizedArguments(args []string) *Alert {
	a.LocArgs = args
	return a
}

func (a *Alert) LaunchImage(image string) *Alert {
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
