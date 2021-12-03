package controllers

import (
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type Register struct {
	Controller
}

func (r *Register) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "auth"
	p.Name = "register"
	p.Title = "Register"
	p.Data = "This is the login page"
	return r.RenderPage(c, p)
}

func (r *Register) Post(c echo.Context) error {
	msg.Set(c, msg.Danger, "Registration is currently disabled.")
	return r.Get(c)
}
