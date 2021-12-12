package controllers

import (
	"goweb/auth"

	"github.com/labstack/echo/v4"
)

type Home struct {
	Controller
}

func (h *Home) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "main"
	p.Name = "home"
	p.Data = "Hello world"
	p.Metatags.Description = "Welcome to the homepage."
	p.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}

	uid, _ := auth.GetUserID(c)
	c.Logger().Infof("logged in user ID: %d", uid)

	return h.RenderPage(c, p)
}
