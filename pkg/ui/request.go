package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"maragu.dev/gomponents"
)

type (
	// Request encapsulates information about the incoming request in order to provide your ui with important and
	// useful information needed for rendering.
	Request struct {
		// Title stores the title of the page.
		Title string

		// Context stores the request context.
		Context echo.Context

		// CurrentPath stores the path of the current request.
		CurrentPath string

		// IsHome stores whether the requested page is the home page.
		IsHome bool

		// IsAuth stores whether the user is authenticated.
		IsAuth bool

		// IsAdmin stores whether the user is an admin.
		IsAdmin bool

		// AuthUser stores the authenticated user.
		AuthUser *ent.User

		// Metatags stores metatag values.
		Metatags struct {
			// Description stores the description metatag value.
			Description string

			// Keywords stores the keywords metatag values.
			Keywords []string
		}

		// CSRF stores the CSRF token for the given request.
		// This will only be populated if the CSRF middleware is in effect for the given request.
		// If this is populated, all forms must include this value otherwise the requests will be rejected.
		CSRF string

		// Htmx stores information provided by HTMX about this request.
		Htmx *htmx.Request

		// Config stores the application configuration.
		// This will only be populated if the Config middleware is installed in the router.
		Config *config.Config
	}

	// LayoutFunc is a callback function intended to render your page node within a given layout.
	// This is handled as a callback to automatically support HTMX requests so that you can respond
	// with only the page content and not the entire layout.
	// See Request.Render().
	LayoutFunc func(*Request, gomponents.Node) gomponents.Node
)

// NewRequest generates a new Request using the Echo context of a given HTTP request.
func NewRequest(ctx echo.Context) *Request {
	p := &Request{
		Context:     ctx,
		CurrentPath: ctx.Request().URL.Path,
		Htmx:        htmx.GetRequest(ctx),
	}

	p.IsHome = p.CurrentPath == "/"

	if csrf := ctx.Get(context.CSRFKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		p.IsAuth = true
		p.AuthUser = u.(*ent.User)
		p.IsAdmin = p.AuthUser.Admin
	}

	if cfg := ctx.Get(context.ConfigKey); cfg != nil {
		p.Config = cfg.(*config.Config)
	}

	return p
}

// Path generates a URL path for a given route name and optional route parameters.
// This will only work if you've supplied names for each of your routes. It's optional to use and helps avoids
// having duplicate, hard-coded paths and parameters all over your application.
func (r *Request) Path(routeName string, routeParams ...any) string {
	return r.Context.Echo().Reverse(routeName, routeParams...)
}

// Url generates an absolute URL for a given route name and optional route parameters.
func (r *Request) Url(routeName string, routeParams ...any) string {
	return r.Config.App.Host + r.Path(routeName, routeParams...)
}

// Render renders a given node, optionally within a given layout based on the HTMX request headers.
// If the request is being made by HTMX and is not boosted, this will automatically only render the node without
// the layout, to support partial rendering.
func (r *Request) Render(layout LayoutFunc, node gomponents.Node) error {
	if r.Htmx.Enabled && !r.Htmx.Boosted {
		return node.Render(r.Context.Response().Writer)
	}

	return layout(r, node).Render(r.Context.Response().Writer)
}
