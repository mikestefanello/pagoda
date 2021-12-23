package routes

import (
	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type About struct {
	controller.Controller
}

func (a *About) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "main"
	p.Name = "about"
	p.Title = "About"
	p.Data = "This is the about page"
	p.Cache.Enabled = false
	p.Cache.Tags = []string{"page_about", "page:list"}

	return a.RenderPage(c, p)
}
