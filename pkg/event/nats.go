package event

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/azimuth3d/woof-service/pkg/schema"
	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	Connection             *nats.Conn
	WoofCreateSubscription *nats.Subscription
	WoofCreateChan         chan WoofCreatedMessage
}

func NewNatsEventStore(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)

	if err != nil {
		return nil, err
	}

	return &NatsEventStore{Connection: nc}, nil
}

func (nes *NatsEventStore) SubscribeWoofCreated() (<-chan WoofCreatedMessage, error) {
	m := WoofCreatedMessage{}
	nes.WoofCreateChan = make(chan WoofCreatedMessage, 64)

	ch := make(chan *nats.Msg, 64)

	var err error

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case msg := <-ch:
				if err := nes.ReadMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				nes.WoofCreateChan <- m
			}
		}
	}()

	return (<-chan WoofCreatedMessage)(nes.WoofCreateChan), nil
}

func (nes *NatsEventStore) OnWoofCreated(f func(WoofCreatedMessage)) (err error) {
	m := WoofCreatedMessage{}
	nes.WoofCreateSubscription, err = nes.Connection.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := nes.ReadMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})

	if err != nil {
		return err
	}
	return nil
}

func (nes *NatsEventStore) Close() {
	if nes.Connection != nil {
		nes.Connection.Close()
	}

	if nes.WoofCreateSubscription != nil {
		if err := nes.WoofCreateSubscription.Unsubscribe(); err != nil {

		}
	}
	close(nes.WoofCreateChan)
}

func (nes *NatsEventStore) PublishWoofCreated(woof schema.Woof) error {
	m := WoofCreatedMessage{woof.ID.String(), woof.Body, time.Unix(int64(woof.CreatedAt.T), 0)}
	data, err := nes.WriteMessage(&m)
	if err != nil {
		return err
	}
	return nes.Connection.Publish(m.Key(), data)
}

func (nes *NatsEventStore) WriteMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)

	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (ens *NatsEventStore) ReadMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
