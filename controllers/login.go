package controllers

import (
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

//func (a *Contact) Post(c echo.Context) error {
//	msg.Set(c, msg.Success, "Thank you for contacting us!")
//	msg.Set(c, msg.Info, "We will respond to you shortly.")
//	return a.Redirect(c, "home")
//}
