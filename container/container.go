package container

import (
	"context"
	"fmt"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"

	"goweb/config"
)

type Container struct {
	Web    *echo.Echo
	Config *config.Config
	Cache  *cache.Cache
	// DB
}

func NewContainer() *Container {
	var c Container

	// Web
	c.Web = echo.New()

	// Configuration
	cfg, err := config.GetConfig()
	if err != nil {
		c.Web.Logger.Error(err)
		c.Web.Logger.Fatal("Failed to load configuration")
	}
	c.Config = &cfg

	// Cache
	cacheClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Config.Cache.Hostname, c.Config.Cache.Port),
		Password: c.Config.Cache.Password,
	})
	if _, err = cacheClient.Ping(context.Background()).Result(); err != nil {
		c.Web.Logger.Error(err)
		c.Web.Logger.Fatal("Failed to connect to cache server")
	}
	cacheStore := store.NewRedis(cacheClient, nil)
	c.Cache = cache.New(cacheStore)

	return &c
}
