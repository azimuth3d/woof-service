package event

import "github.com/azimuth3d/woof-service/pkg/schema"

type EventStore interface {
	Close()
	PublishWoofCreate(woof schema.Woof) error
	SubscribeWoofCreated(<-chan WoofCreatedMessage) error
	OnWoofCreated(f func(WoofCreatedMessage)) error
}
