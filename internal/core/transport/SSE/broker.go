package core_sse

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type Broker struct {
	notifier       chan Event
	newClients     chan client
	closingClients chan chan Event
	clients        map[chan Event]string
	mu             sync.Mutex
}

type client struct {
	ch      chan Event
	boardID string
}

func NewBroker() *Broker {
	return &Broker{
		notifier:       make(chan Event, 1),
		newClients:     make(chan client),
		closingClients: make(chan chan Event),
		clients:        make(map[chan Event]string),
	}
}

func (b *Broker) Listen() {
	for {
		select {
		case c := <-b.newClients:
			b.addClient(c)
			log.Println("Клиент подключился. Всего:", len(b.clients))

		case ch := <-b.closingClients:
			b.removeClient(ch)
			log.Println("Клиент отключился. Всего:", len(b.clients))

		case event := <-b.notifier:
			b.mu.Lock()
			for ch, boardID := range b.clients {
				if boardID != event.BoardID {
					continue
				}
				select {
				case ch <- event:
				default:
					log.Println("Буфер клиента заполнен, сообщение пропущено")
				}
			}
			b.mu.Unlock()
		}
	}
}

func (b *Broker) addClient(c client) {
	b.mu.Lock()
	b.clients[c.ch] = c.boardID
	b.mu.Unlock()
}

func (b *Broker) removeClient(ch chan Event) {
	b.mu.Lock()
	if _, exists := b.clients[ch]; exists {
		delete(b.clients, ch)
		close(ch)
	}
	b.mu.Unlock()
}

func (b *Broker) Publish(boardID, eventType string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event %s: %w", eventType, err)
	}

	b.notifier <- Event{BoardID: boardID, Type: eventType, Data: data}
	return nil
}
