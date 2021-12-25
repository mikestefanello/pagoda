package routes

import (
	"net/http"

	"goweb/config"
	"goweb/controller"
	"goweb/middleware"
	"goweb/services"

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
	err := Error{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Example routes
	navRoutes(c, g, ctr)
	userRoutes(c, g, ctr)
}

func navRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	home := Home{Controller: ctr}
	g.GET("/", home.Get).Name = "home"

	about := About{Controller: ctr}
	g.GET("/about", about.Get).Name = "about"

	contact := Contact{Controller: ctr}
	g.GET("/contact", contact.Get).Name = "contact"
	g.POST("/contact", contact.Post).Name = "contact.post"
}

func userRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	logout := Logout{Controller: ctr}
	g.GET("/logout", logout.Get, middleware.RequireAuthentication()).Name = "logout"

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	login := Login{Controller: ctr}
	noAuth.GET("/login", login.Get).Name = "login"
	noAuth.POST("/login", login.Post).Name = "login.post"

	register := Register{Controller: ctr}
	noAuth.GET("/register", register.Get).Name = "register"
	noAuth.POST("/register", register.Post).Name = "register.post"

	forgot := ForgotPassword{Controller: ctr}
	noAuth.GET("/password", forgot.Get).Name = "forgot_password"
	noAuth.POST("/password", forgot.Post).Name = "forgot_password.post"

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(c.ORM),
		middleware.LoadValidPasswordToken(c.Auth),
	)
	reset := ResetPassword{Controller: ctr}
	resetGroup.GET("/token/:user/:password_token", reset.Get).Name = "reset_password"
	resetGroup.POST("/token/:user/:password_token", reset.Post).Name = "reset_password.post"
}
