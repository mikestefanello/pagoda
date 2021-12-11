package controllers

import (
	"net/http"

	"goweb/config"
	"goweb/middleware"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"goweb/container"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// TODO: This is doing more than building the router

func BuildRouter(c *container.Container) {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Middleware
	g := c.Web.Group("",
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
		middleware.PageCache(c.Cache),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	// Base controller
	ctr := NewController(c)

	// Error handler
	err := Error{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Validator
	c.Web.Validator = &Validator{validator: validator.New()}

	// Routes
	navRoutes(g, ctr)
	userRoutes(g, ctr)
}

func navRoutes(g *echo.Group, ctr Controller) {
	home := Home{Controller: ctr}
	g.GET("/", home.Get).Name = "home"

	about := About{Controller: ctr}
	g.GET("/about", about.Get).Name = "about"

	contact := Contact{Controller: ctr}
	g.GET("/contact", contact.Get).Name = "contact"
	g.POST("/contact", contact.Post).Name = "contact.post"
}

func userRoutes(g *echo.Group, ctr Controller) {
	login := Login{Controller: ctr}
	g.GET("/user/login", login.Get).Name = "login"
	g.POST("/user/login", login.Post).Name = "login.post"

	register := Register{Controller: ctr}
	g.GET("/user/register", register.Get).Name = "register"
	g.POST("/user/register", register.Post).Name = "register.post"
}
