package routes

import (
	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type About struct {
	controller.Controller
}

func (c *About) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "about"
	page.Title = "About"
	page.Data = "The data field can take in anything you want to send to the templates"
	page.Cache.Enabled = false
	page.Cache.Tags = []string{"page_about", "page:list"}

	return c.RenderPage(ctx, page)
}
