package main

import (
	"sync"
	"time"
)

type (
	ServerState struct {
		Users map[string]User
		mutex sync.RWMutex
	}
	User struct {
		LastPoll time.Time
		Name     string
	}
)

func NewState() *ServerState {
	return &ServerState{
		Users: map[string]User{},
	}
}
