package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	inertia "github.com/romsar/gonertia/v2"
)

type Dashboard struct {
	Inertia *inertia.Inertia
}

func init() {
	Register(new(Dashboard))
}

func (h *Dashboard) Init(c *services.Container) error {
	h.Inertia = c.Inertia
	return nil
}

func (h *Dashboard) Routes(g *echo.Group) {
	g.GET("/dashboard", h.Page).Name = routenames.Dashboard
}

func (h *Dashboard) Page(ctx echo.Context) error {
	err := h.Inertia.Render(
		ctx.Response().Writer,
		ctx.Request(),
		"Dashboard",
		inertia.Props{
			"title": "Welcome to the Dashboard",
		},
	)
	if err != nil {
		handleServerErr(ctx.Response().Writer, err)
		return err
	}

	return nil
}
