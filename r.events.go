package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *State) Events(c echo.Context) error {
	cookie, err := c.Cookie(authCookie)
	if err != nil {
		return s.Login(c)
	}

	userId := cookie.Value

	clientId := uuid.New().String()
	fmt.Printf("SSE client connected, clientId: %v\n", clientId)
	ch := s.clients.AddClient(userId, clientId)
	defer s.clients.RemoveClient(clientId)

	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	// go s.clients.SendAll(Event{key: "message", value: userId})
	// go s.clients.SendAll(Event{key: "test", value: clientId})

	for {
		select {
		case <-c.Request().Context().Done():
			fmt.Printf("SSE client disconnected, clientId: %v\n", clientId)
			return nil
		case event := <-ch:
			fmt.Fprintf(c.Response(), "event: %s\ndata: <div>%s</div>\n\n", event.key, event.value)
			c.Response().Flush()
		}
	}
}
