package main

type (
	State struct {
		users   UserState
		clients ClientState
	}
)

func NewState() *State {
	return &State{
		users:   *NewUserState(),
		clients: *NewClientState(),
	}
}
