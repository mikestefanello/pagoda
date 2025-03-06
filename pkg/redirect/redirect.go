package redirect

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
)

// Redirect is a helper to perform HTTP redirects.
type Redirect struct {
	ctx         echo.Context
	url         string
	routeName   string
	routeParams []any
	status      int
	query       url.Values
}

// New initializes a new Redirect
func New(ctx echo.Context) *Redirect {
	return &Redirect{
		ctx:    ctx,
		status: http.StatusTemporaryRedirect,
	}
}

// Route sets the route name to redirect to.
// Use either this or URL()
func (r *Redirect) Route(name string) *Redirect {
	r.routeName = name
	return r
}

// Params sets the route params
func (r *Redirect) Params(params ...any) *Redirect {
	r.routeParams = params
	return r
}

// StatusCode sets the HTTP status code which defaults to http.StatusFound.
// Does not apply to HTMX redirects.
func (r *Redirect) StatusCode(code int) *Redirect {
	r.status = code
	return r
}

// Query sets a URL query
func (r *Redirect) Query(query url.Values) *Redirect {
	r.query = query
	return r
}

// URL sets the URL to redirect to
// Use either this or Route()
func (r *Redirect) URL(url string) *Redirect {
	r.url = url
	return r
}

// Go performs the redirect
// If the request is HTMX boosted, an HTMX redirect will be performed instead of an HTTP redirect
func (r *Redirect) Go() error {
	if r.routeName == "" && r.url == "" {
		return errors.New("no redirect provided")
	}

	var dest string
	if r.url != "" {
		dest = r.url
	} else {
		dest = r.ctx.Echo().Reverse(r.routeName, r.routeParams...)
	}

	if len(r.query) > 0 {
		dest = fmt.Sprintf("%s?%s", dest, r.query.Encode())
	}

	if htmx.GetRequest(r.ctx).Boosted {
		htmx.Response{
			Redirect: dest,
		}.Apply(r.ctx)

		return nil
	} else {
		return r.ctx.Redirect(r.status, dest)
	}
}
