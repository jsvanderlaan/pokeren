package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *State) Logout(c echo.Context) error {
	cookie, err := c.Cookie(authCookie)
	if err == nil {
		guid := cookie.Value
		s.users.DeleteUser(guid)
	}

	// Remove cookie
	cookie = &http.Cookie{
		Name:    authCookie,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}
	c.SetCookie(cookie)

	return s.Login(c)
}
