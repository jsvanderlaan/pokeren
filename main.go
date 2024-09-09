package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const cookieName string = "user-id"

type PageState struct {
	Login bool
}

func main() {
	e := echo.New()
	s := NewState()
	e.Use(middleware.Logger())
	// e.Use(s.Authenticate)

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", PageState{Login: !hasCookie(c)})
	})
	e.POST("/login", func(c echo.Context) error {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		fmt.Println(c.Request())

		return c.Render(200, "main", nil)
	})

	e.Logger.Fatal(e.Start(":42069"))
}

func hasCookie(c echo.Context) bool {
	_, err := c.Cookie(cookieName)
	return err == nil
}
