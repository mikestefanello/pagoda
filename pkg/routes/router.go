package routes

import (
	"net/http"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non static file route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.Logger(),
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.Cache),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	// Base controller
	ctr := controller.NewController(c)

	// Error handler
	err := errorHandler{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Example routes
	navRoutes(c, g, ctr)
	userRoutes(c, g, ctr)
}

func navRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	home := home{Controller: ctr}
	g.GET("/", home.Get).Name = "home"

	search := search{Controller: ctr}
	g.GET("/search", search.Get).Name = "search"

	about := about{Controller: ctr}
	g.GET("/about", about.Get).Name = "about"

	contact := contact{Controller: ctr}
	g.GET("/contact", contact.Get).Name = "contact"
	g.POST("/contact", contact.Post).Name = "contact.post"
}

func userRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	logout := logout{Controller: ctr}
	g.GET("/logout", logout.Get, middleware.RequireAuthentication()).Name = "logout"

	verifyEmail := verifyEmail{Controller: ctr}
	g.GET("/email/verify/:token", verifyEmail.Get).Name = "verify_email"

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	login := login{Controller: ctr}
	noAuth.GET("/login", login.Get).Name = "login"
	noAuth.POST("/login", login.Post).Name = "login.post"

	register := register{Controller: ctr}
	noAuth.GET("/register", register.Get).Name = "register"
	noAuth.POST("/register", register.Post).Name = "register.post"

	forgot := forgotPassword{Controller: ctr}
	noAuth.GET("/password", forgot.Get).Name = "forgot_password"
	noAuth.POST("/password", forgot.Post).Name = "forgot_password.post"

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(c.ORM),
		middleware.LoadValidPasswordToken(c.Auth),
	)
	reset := resetPassword{Controller: ctr}
	resetGroup.GET("/token/:user/:password_token/:token", reset.Get).Name = "reset_password"
	resetGroup.POST("/token/:user/:password_token/:token", reset.Post).Name = "reset_password.post"
}
