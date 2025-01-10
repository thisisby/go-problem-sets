package main

import (
	"log"
	"sync"
)

func main() {
	m := PubSubManager{
		subscribers: make(map[string][]chan string),
	}

	msgChan1 := m.Subscribe("id1")
	msgChan2 := m.Subscribe("id1")
	msgChan3 := m.Subscribe("id2")

	m.Publish("id1", "hello world")
	m.Publish("id1", "hi")
	m.Publish("id1", "test")
	m.Publish("id2", "testing!!!")

	log.Println(<-msgChan1, <-msgChan1, <-msgChan1)
	log.Println(<-msgChan2, <-msgChan2, <-msgChan2)
	log.Println(<-msgChan3)
}

type PubSubManager struct {
	subscribers map[string][]chan string
	mu          sync.Mutex
}

func (p *PubSubManager) Subscribe(topic string) chan string {
	p.mu.Lock()
	defer p.mu.Unlock()

	msgChan := make(chan string)

	p.subscribers[topic] = append(p.subscribers[topic], msgChan)

	return msgChan
}

func (p *PubSubManager) Publish(topic string, message string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, ch := range p.subscribers[topic] {
		go func(c chan string) {
			c <- message
		}(ch)
	}
}
