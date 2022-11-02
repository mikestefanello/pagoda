package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// CachedPageGroup stores the cache group for cached pages
const CachedPageGroup = "page"

// CachedPage is what is used to store a rendered Page in the cache
type CachedPage struct {
	// URL stores the URL of the requested page
	URL string

	// HTML stores the complete HTML of the rendered Page
	HTML []byte

	// StatusCode stores the HTTP status code
	StatusCode int

	// Headers stores the HTTP headers
	Headers map[string]string
}

// ServeCachedPage attempts to load a page from the cache by matching on the complete request URL
// If a page is cached for the requested URL, it will be served here and the request terminated.
// Any request made by an authenticated user or that is not a GET will be skipped.
func ServeCachedPage(ch *services.CacheClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip non GET requests
			if c.Request().Method != http.MethodGet {
				return next(c)
			}

			// Skip if the user is authenticated
			if c.Get(context.AuthenticatedUserKey) != nil {
				return next(c)
			}

			// Attempt to load from cache
			res, err := ch.
				Get().
				Group(CachedPageGroup).
				Key(c.Request().URL.String()).
				Type(new(CachedPage)).
				Fetch(c.Request().Context())

			if err != nil {
				switch {
				case err == redis.Nil:
					c.Logger().Info("no cached page found")
				case context.IsCanceledError(err):
					return nil
				default:
					c.Logger().Errorf("failed getting cached page: %v", err)
				}

				return next(c)
			}

			page, ok := res.(*CachedPage)
			if !ok {
				c.Logger().Errorf("failed casting cached page")
				return next(c)
			}

			// Set any headers
			if page.Headers != nil {
				for k, v := range page.Headers {
					c.Response().Header().Set(k, v)
				}
			}

			c.Logger().Info("serving cached page")

			return c.HTMLBlob(page.StatusCode, page.HTML)
		}
	}
}

// CacheControl sets a Cache-Control header with a given max age
func CacheControl(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v := "no-cache, no-store"
			if maxAge > 0 {
				v = fmt.Sprintf("public, max-age=%.0f", maxAge.Seconds())
			}
			c.Response().Header().Set("Cache-Control", v)
			return next(c)
		}
	}
}
