package routes

import (
	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type Home struct {
	controller.Controller
}

func (h *Home) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "main"
	p.Name = "home"
	p.Data = "Hello world"
	p.Metatags.Description = "Welcome to the homepage."
	p.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}

	return h.RenderPage(c, p)
}
