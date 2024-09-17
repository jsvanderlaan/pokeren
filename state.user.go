package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

const dataFile = "users.json"

type (
	UserState struct {
		mutex sync.RWMutex
	}
	User struct {
		LastPoll time.Time
		Username string
		Guid     string
	}
)

func NewUserState() *UserState {
	return &UserState{}
}

func NewUser(guid string, username string) *User {
	return &User{
		Username: username,
		Guid:     guid,
		LastPoll: time.Now(),
	}
}

func (s *UserState) ReadUsers() ([]User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserState) GetUser(guid string) (User, error) {
	users, err := s.ReadUsers()
	if err != nil {
		return User{}, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, user := range users {
		if user.Guid == guid {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}

func (s *UserState) UpdateUser(guid string) ([]User, error) {
	users, err := s.ReadUsers()
	if err != nil {
		return nil, err
	}
	newUsers := []User{}

	for _, user := range users {
		if user.Guid == guid {
			user.LastPoll = time.Now()
			newUsers = append(newUsers, user)
			continue
		}

		if user.LastPoll.Before(time.Now().Add(-1 * time.Second * 30)) {
			continue
		}

		newUsers = append(newUsers, user)

	}
	s.WriteUsers(newUsers)

	return newUsers, nil
}

func (s *UserState) DeleteUser(guid string) error {
	users, err := s.ReadUsers()
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, user := range users {
		if user.Guid == guid {
			users = append(users[:i], users[i+1:]...)
			file, err := os.Create(dataFile)
			if err != nil {
				return err
			}
			defer file.Close()
			json.NewEncoder(file).Encode(users)
			return nil
		}
	}

	return nil
}

func (s *UserState) WriteUser(user User) error {
	users, err := s.ReadUsers()
	if err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	users = append(users, user)

	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func (s *UserState) WriteUsers(users []User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}
