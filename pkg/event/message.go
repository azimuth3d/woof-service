package event

import "time"

type Message interface {
	Key() string
}

type WoofCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (m *WoofCreatedMessage) Key() string {
	return "WoofCreate"
}
