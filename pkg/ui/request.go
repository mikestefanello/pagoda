package ui

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"maragu.dev/gomponents"
)

type layoutFunc func(*request, gomponents.Node) gomponents.Node

type request struct {
	// AppName stores the name of the application.
	// If omitted, the configuration value will be used.
	AppName string

	// Title stores the title of the page
	Title string

	// Context stores the request context
	Context echo.Context

	// Path stores the path of the current request
	Path string

	// IsHome stores whether the requested page is the home page or not
	IsHome bool

	// IsAuth stores whether the user is authenticated
	IsAuth bool

	// AuthUser stores the authenticated user
	AuthUser *ent.User

	// Metatags stores metatag values
	Metatags struct {
		// Description stores the description metatag value
		Description string

		// Keywords stores the keywords metatag values
		Keywords []string
	}

	// CSRF stores the CSRF token for the given request.
	// This will only be populated if the CSRF middleware is in effect for the given request.
	// If this is populated, all forms must include this value otherwise the requests will be rejected.
	CSRF string

	Htmx htmx.Request
}

func newRequest(ctx echo.Context) *request {
	p := &request{
		Context: ctx,
		Path:    ctx.Request().URL.Path,
		Htmx:    htmx.GetRequest(ctx),
	}

	p.IsHome = p.Path == "/"

	if csrf := ctx.Get(echomw.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		p.IsAuth = true
		p.AuthUser = u.(*ent.User)
	}

	return p
}

func (r *request) path(routeName string, routeParams ...string) string {
	return r.Context.Echo().Reverse(routeName, routeParams)
}

func (r *request) render(layout layoutFunc, node gomponents.Node) error {
	if r.Htmx.Enabled && !r.Htmx.Boosted {
		return node.Render(r.Context.Response().Writer)
	}

	return layout(r, node).Render(r.Context.Response().Writer)
}
