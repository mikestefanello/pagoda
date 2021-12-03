package router

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"goweb/container"
	"goweb/controllers"
)

const StaticDir = "static"

func BuildRouter(c *container.Container) {
	// Middleware
	c.Web.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	c.Web.Use(middleware.RequestID())
	c.Web.Use(middleware.Recover())
	c.Web.Use(middleware.Gzip())
	c.Web.Use(middleware.Logger())
	// TODO: needs cache control headers
	c.Web.Use(middleware.Static(StaticDir))
	c.Web.Use(session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))))
	c.Web.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

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
	// TODO
}
