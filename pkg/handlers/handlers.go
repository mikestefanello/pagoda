package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
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

// redirect redirects to a given route by name with optional route parameters
func redirect(ctx echo.Context, routeName string, routeParams ...any) error {
	return doRedirect(ctx, ctx.Echo().Reverse(routeName, routeParams...))
}

// redirectWithQuery redirects to a given route by name with query parameters and optional route parameters
func redirectWithQuery(ctx echo.Context, query url.Values, routeName string, routeParams ...any) error {
	dest := fmt.Sprintf("%s?%s", ctx.Echo().Reverse(routeName, routeParams...), query.Encode())
	return doRedirect(ctx, dest)
}

// doRedirect performs a redirect to a given URL
func doRedirect(ctx echo.Context, url string) error {
	if htmx.GetRequest(ctx).Boosted {
		htmx.Response{
			Redirect: url,
		}.Apply(ctx)

		return nil
	} else {
		return ctx.Redirect(http.StatusFound, url)
	}
}

// fail is a helper to fail a request by returning a 500 error and logging the error
func fail(err error, log string) error {
	// The error handler will handle logging
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", log, err))
}
