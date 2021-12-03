package controllers

import (
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type Login struct {
	Controller
}

func (l *Login) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "auth"
	p.Name = "login"
	p.Title = "Log in"
	p.Data = "This is the login page"
	return l.RenderPage(c, p)
}

func (l *Login) Post(c echo.Context) error {
	msg.Set(c, msg.Danger, "Invalid credentials. Please try again.")
	return l.Get(c)
}
