package handlers

import (
	"github.com/labstack/echo/v4"
)

var handlers []Routable

type Routable interface {
	Routes(g *echo.Group)
}

func Register(r Routable) {
	handlers = append(handlers, r)
}

func GetHandlers() []Routable {
	return handlers
}
