// Copyright 2012-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event

import (
	"log"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestDefaultConnection(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection(t)
	nc.Close()
}

func TestConnectionStatus(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	nc := NewDefaultConnection(t)
	defer nc.Close()

	if nc.Status() != nats.CONNECTED || nc.Status().String() != "CONNECTED" {
		t.Fatal("Should have status set to CONNECTED")
	}

	if !nc.IsConnected() {
		t.Fatal("Should have status set to CONNECTED")
	}
	nc.Close()
	if nc.Status() != nats.CLOSED || nc.Status().String() != "CLOSED" {
		t.Fatal("Should have status set to CLOSED")
	}
	if !nc.IsClosed() {
		t.Fatal("Should have status set to CLOSED")
	}
}

func TestNewNatsEventStore(t *testing.T) {
	s := RunDefaultServer()
	defer s.Shutdown()

	// nc := NewDefaultConnection(t)

	nes, err := NewNatsEventStore("nats://127.0.0.1:4222")

	if err != nil {
		log.Fatal(err)
	}

	t.Log(nes.Connection.Status())

	if nes.Connection == nil {
		t.Fatal("Should have connection")
	}
}
