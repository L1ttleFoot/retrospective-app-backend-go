package core_sse

import (
	"log"
	"sync"
)

type Broker struct {
	notifier       chan []byte
	newClients     chan chan []byte
	closingClients chan chan []byte
	clients        map[chan []byte]struct{}
	mu             sync.Mutex
}

func NewBroker() *Broker {
	b := &Broker{
		notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]struct{}),
	}

	return b
}

func (b *Broker) Listen() {
	for {
		select {
		case client := <-b.newClients:
			b.addClient(client)
			log.Println("Клиент подключился. Всего:", len(b.clients))

		case client := <-b.closingClients:
			b.removeClient(client)
			log.Println("Клиент отключился. Всего:", len(b.clients))

		case event := <-b.notifier:
			b.mu.Lock()
			for clientChan := range b.clients {
				select {
				case clientChan <- event:
				default:
					log.Println("Буфер клиента заполнен, сообщение пропущено")
				}
			}
			b.mu.Unlock()
		}
	}
}

func (b *Broker) addClient(ch chan []byte) {
	b.mu.Lock()
	b.clients[ch] = struct{}{}
	b.mu.Unlock()
}

func (b *Broker) removeClient(ch chan []byte) {
	b.mu.Lock()
	if _, exists := b.clients[ch]; exists {
		delete(b.clients, ch)
		close(ch)
	}
	b.mu.Unlock()
}

func (b *Broker) Publish(msg []byte) {
	b.notifier <- msg
}
