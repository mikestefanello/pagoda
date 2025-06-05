package handlers

import (
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"
	files "github.com/mikestefanello/pagoda/public"
)

// BuildRouter builds the router.
func BuildRouter(c *services.Container) error {
	// Force HTTPS, if enabled.
	if c.Config.HTTP.TLS.Enabled {
		c.Web.Use(echomw.HTTPSRedirect())
	}

	// Serve public files with cache control.
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.PublicFile)).
		Static("files", "public/files")

	// Serve static files.
	// ui.StaticFile() should be used in ui components to append a cache key to the URL to break cache
	// after each server reboot.
	c.Web.Group(
		"",
		echomw.GzipWithConfig(echomw.GzipConfig{
			Skipper: func(c echo.Context) bool {
				for _, ext := range []string{
					".js",
					".css",
				} {
					if strings.HasSuffix(c.Request().URL.Path, ext) {
						return false
					}
				}
				return true
			},
		}),
	).StaticFS("static", echo.MustSubFS(files.Static, "static"))

	// Non-static file route group.
	g := c.Web.Group("")

	// Create a cookie store for session data.
	cookieStore := sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.SameSite = http.SameSiteStrictMode

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		middleware.Config(c.Config),
		middleware.Session(cookieStore),
		middleware.LoadAuthenticatedUser(c.Auth),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup:    "form:csrf",
			CookieHTTPOnly: true,
			CookieSameSite: http.SameSiteStrictMode,
			ContextKey:     context.CSRFKey,
		}),
	)

	// Error handler.
	c.Web.HTTPErrorHandler = new(Error).Page

	// Initialize and register all handlers.
	for _, h := range GetHandlers() {
		if err := h.Init(c); err != nil {
			return err
		}

		h.Routes(g)
	}

	return nil
}
