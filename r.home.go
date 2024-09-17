package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *State) Home(c echo.Context) error {
	cookie, err := c.Cookie(authCookie)
	if err != nil {
		return s.Login(c)
	}

	guid := cookie.Value
	user, err := s.users.GetUser(guid)
	if err != nil {
		return s.Logout(c)
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"guid":     guid,
		"username": user.Username,
	})
}
