package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const authCookie string = "guid"

func main() {
	e := echo.New()
	s := NewState()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.GET("/", Index)
	e.GET("/home", s.Home)
	// e.GET("/poll", s.PollGet)
	e.POST("/login", s.LoginPost)
	e.GET("/logout", s.Logout)

	e.GET("/events", s.Events)

	e.Logger.Fatal(e.Start(":42069"))
}
