package router

import (
	"net/http"

	"goweb/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"goweb/container"
	"goweb/controllers"
)

const StaticDir = "static"

func BuildRouter(c *container.Container) {
	// Middleware
	c.Web.Use(echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	c.Web.Use(echomw.RequestID())
	c.Web.Use(echomw.Recover())
	c.Web.Use(echomw.Gzip())
	c.Web.Use(echomw.Logger())
	c.Web.Use(session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))))
	c.Web.Use(echomw.CSRFWithConfig(echomw.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(15552000)).
		Static("/", StaticDir)

	// Base controller
	ctr := controllers.NewController(c)

	// Error handler
	err := controllers.Error{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Routes
	navRoutes(c.Web, ctr)
	userRoutes(c.Web, ctr)
}

func navRoutes(e *echo.Echo, ctr controllers.Controller) {
	home := controllers.Home{Controller: ctr}
	e.GET("/", home.Get).Name = "home"

	about := controllers.About{Controller: ctr}
	e.GET("/about", about.Get).Name = "about"

	contact := controllers.Contact{Controller: ctr}
	e.GET("/contact", contact.Get).Name = "contact"
	e.POST("/contact", contact.Post).Name = "contact.post"
}

func userRoutes(e *echo.Echo, ctr controllers.Controller) {
	login := controllers.Login{Controller: ctr}
	e.GET("/user/login", login.Get).Name = "login"
}
