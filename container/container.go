package container

import (
	"github.com/labstack/echo/v4"

	"goweb/config"
)

type Container struct {
	Web    *echo.Echo
	Config *config.Config
	// Cache
	// DB
}

func NewContainer() *Container {
	var c Container

	// Web
	c.Web = echo.New()

	// Configuration
	cfg, err := config.GetConfig()
	if err != nil {
		c.Web.Logger.Fatal("Failed to load configuration")
		panic(err)
	}
	c.Config = &cfg

	return &c
}
