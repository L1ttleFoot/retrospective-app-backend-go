package core_sse

import (
	"testing"
	"time"
)

func TestBroker_PublishFiltersByBoard(t *testing.T) {
	b := NewBroker()
	go b.Listen()

	ch1 := make(chan Event, 1)
	ch2 := make(chan Event, 1)
	b.newClients <- client{ch: ch1, boardID: "board-1"}
	b.newClients <- client{ch: ch2, boardID: "board-2"}

	if err := b.Publish("board-1", "message.created", map[string]string{"id": "m1"}); err != nil {
		t.Fatalf("publish: %v", err)
	}

	select {
	case ev := <-ch1:
		if ev.Type != "message.created" || string(ev.Data) != `{"id":"m1"}` {
			t.Fatalf("unexpected event: %+v", ev)
		}
	case <-time.After(time.Second):
		t.Fatal("клиент board-1 не получил событие")
	}

	select {
	case ev := <-ch2:
		t.Fatalf("клиент board-2 получил событие board-1: %+v", ev)
	case <-time.After(50 * time.Millisecond):
	}
}
