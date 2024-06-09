package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
)

var handlers []Handler

// Handler handles one or more HTTP routes
type Handler interface {
	// Routes allows for self-registration of HTTP routes on the router
	Routes(g *echo.Group)

	// Init provides the service container to initialize
	Init(*services.Container) error
}

// Register registers a handler
func Register(h Handler) {
	handlers = append(handlers, h)
}

// GetHandlers returns all handlers
func GetHandlers() []Handler {
	return handlers
}
