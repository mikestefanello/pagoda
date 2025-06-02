package handlers

import (
	"net/http"

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
	// Public files with proper cache control.
	// ui.PublicFile() should be used in ui components to append a cache key to the URL to break cache
	// after each server restart.
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.PublicFile)).
		Static("files", "public/files")

	// Static files with proper cache control.
	// ui.StaticFile() should be used in ui components to append a cache key to the URL to break cache
	// after each server restart.
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.PublicFile)).
		StaticFS("static", echo.MustSubFS(files.Static, "static"))

	// TODO is cache control needed for ^?
	// TODO separate cache control? ^

	// Non-static file route group.
	g := c.Web.Group("")

	// Force HTTPS, if enabled.
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	// Create a cookie store for session data.
	cookieStore := sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = true
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
			CookieSecure:   true,
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
