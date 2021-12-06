package controllers

import (
	"github.com/labstack/echo/v4"
)

type About struct {
	Controller
}

func (a *About) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "main"
	p.Name = "about"
	p.Title = "About"
	p.Data = "This is the about page"
	p.Cache.Enabled = true
	p.Cache.Tags = []string{"page_about", "page:list"}

	return a.RenderPage(c, p)
}
