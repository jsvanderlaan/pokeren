package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *State) Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func (s *State) LoginPost(c echo.Context) error {
	username := c.FormValue("username")
	guid := uuid.New().String()

	c.SetCookie(&http.Cookie{
		Name:  authCookie,
		Value: guid,
		Path:  "/",
	})
	s.users.WriteUser(*NewUser(guid, username))
	return c.Redirect(http.StatusSeeOther, "/")
}
