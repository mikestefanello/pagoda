package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"

	libstore "github.com/eko/gocache/lib/v4/store"
	"github.com/labstack/echo/v4"
)

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
				Group(services.CachedPageGroup).
				Key(c.Request().URL.String()).
				Type(new(services.CachedPage)).
				Fetch(c.Request().Context())

			if err != nil {
				switch {
				case errors.Is(err, &libstore.NotFound{}),
					context.IsCanceledError(err):
					return nil
				default:
					log.Ctx(c).Error("failed getting cached page",
						"error", err,
					)
				}

				return next(c)
			}

			page, ok := res.(*services.CachedPage)
			if !ok {
				log.Ctx(c).Error("failed casting cached page")
				return next(c)
			}

			// Set any headers
			if page.Headers != nil {
				for k, v := range page.Headers {
					c.Response().Header().Set(k, v)
				}
			}

			log.Ctx(c).Debug("serving cached page")

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
