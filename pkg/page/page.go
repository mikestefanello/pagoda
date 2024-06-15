package page

import (
	"html/template"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/templates"

	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// Page consists of all data that will be used to render a page response for a given route.
// While it's not required for a handler to render a Page on a route, this is the common data
// object that will be passed to the templates, making it easy for all handlers to share
// functionality both on the back and frontend. The Page can be expanded to include anything else
// your app wants to support.
// Methods on this page also then become available in the templates, which can be more useful than
// the funcmap if your methods require data stored in the page, such as the context.
type Page struct {
	// AppName stores the name of the application.
	// If omitted, the configuration value will be used.
	AppName string

	// Title stores the title of the page
	Title string

	// Context stores the request context
	Context echo.Context

	// Path stores the path of the current request
	Path string

	// URL stores the URL of the current request
	URL string

	// Data stores whatever additional data that needs to be passed to the templates.
	// This is what the handler uses to pass the content of the page.
	Data any

	// Form stores a struct that represents a form on the page.
	// This should be a struct with fields for each form field, using both "form" and "validate" tags
	// It should also contain form.FormSubmission if you wish to have validation
	// messages and markup presented to the user
	Form any

	// Layout stores the name of the layout base template file which will be used when the page is rendered.
	// This should match a template file located within the layouts directory inside the templates directory.
	// The template extension should not be included in this value.
	Layout templates.Layout

	// Name stores the name of the page as well as the name of the template file which will be used to render
	// the content portion of the layout template.
	// This should match a template file located within the pages directory inside the templates directory.
	// The template extension should not be included in this value.
	Name templates.Page

	// IsHome stores whether the requested page is the home page or not
	IsHome bool

	// IsAuth stores whether the user is authenticated
	IsAuth bool

	// AuthUser stores the authenticated user
	AuthUser *ent.User

	// StatusCode stores the HTTP status code that will be returned
	StatusCode int

	// Metatags stores metatag values
	Metatags struct {
		// Description stores the description metatag value
		Description string

		// Keywords stores the keywords metatag values
		Keywords []string
	}

	// Pager stores a pager which can be used to page lists of results
	Pager Pager

	// CSRF stores the CSRF token for the given request.
	// This will only be populated if the CSRF middleware is in effect for the given request.
	// If this is populated, all forms must include this value otherwise the requests will be rejected.
	CSRF string

	// Headers stores a list of HTTP headers and values to be set on the response
	Headers map[string]string

	// RequestID stores the ID of the given request.
	// This will only be populated if the request ID middleware is in effect for the given request.
	RequestID string

	// HTMX provides the ability to interact with the HTMX library
	HTMX struct {
		// Request contains the information provided by HTMX about the current request
		Request htmx.Request

		// Response contains values to pass back to HTMX
		Response *htmx.Response
	}

	// Cache stores values for caching the response of this page
	Cache struct {
		// Enabled dictates if the response of this page should be cached.
		// Cached responses are served via middleware.
		Enabled bool

		// Expiration stores the amount of time that the cache entry should live for before expiring.
		// If omitted, the configuration value will be used.
		Expiration time.Duration

		// Tags stores a list of tags to apply to the cache entry.
		// These are useful when invalidating cache for dynamic events such as entity operations.
		Tags []string
	}
}

// New creates and initiatizes a new Page for a given request context
func New(ctx echo.Context) Page {
	p := Page{
		Context:    ctx,
		Path:       ctx.Request().URL.Path,
		URL:        ctx.Request().URL.String(),
		StatusCode: http.StatusOK,
		Pager:      NewPager(ctx, DefaultItemsPerPage),
		Headers:    make(map[string]string),
		RequestID:  ctx.Response().Header().Get(echo.HeaderXRequestID),
	}

	p.IsHome = p.Path == "/"

	if csrf := ctx.Get(echomw.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		p.IsAuth = true
		p.AuthUser = u.(*ent.User)
	}

	p.HTMX.Request = htmx.GetRequest(ctx)

	return p
}

// GetMessages gets all flash messages for a given type.
// This allows for easy access to flash messages from the templates.
func (p Page) GetMessages(typ msg.Type) []template.HTML {
	strs := msg.Get(p.Context, typ)
	ret := make([]template.HTML, len(strs))
	for k, v := range strs {
		ret[k] = template.HTML(v)
	}
	return ret
}
