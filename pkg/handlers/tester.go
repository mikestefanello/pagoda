package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/services"
)

func init() {
	Register(new(tester))
}

type tester struct {
	config *config.Config
	controller.Controller
}

func (t *tester) Init(c *services.Container) error {
	t.config = c.Config
	t.Controller = controller.NewController(c)
	return nil
}

func (t *tester) Routes(g *echo.Group) {
	g.GET("/tester", func(c echo.Context) error {
		return c.String(200, "tester "+t.config.App.Name)
	})
}
