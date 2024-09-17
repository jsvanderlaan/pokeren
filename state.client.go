package main

import (
	"sync"
)

type (
	ClientState struct {
		clients map[string]Client
		mutex   sync.RWMutex
	}
	Client struct {
		events chan Event
		userId string
	}
	Event struct {
		key   string
		value string
	}
)

func NewClientState() *ClientState {
	return &ClientState{
		clients: map[string]Client{},
	}
}

func (s *ClientState) SendAll(event Event) {
	for _, client := range s.clients {
		client.events <- event
	}
}

func (s *ClientState) AddClient(userId string, clientId string) chan Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	ch := make(chan Event)
	s.clients[clientId] = Client{
		events: ch,
		userId: userId,
	}
	return ch
}

func (s *ClientState) RemoveClient(clientId string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clients, clientId)
}
