package container

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"

	"goweb/config"
	"goweb/ent"
)

type Container struct {
	Web      *echo.Echo
	Config   *config.Config
	Cache    *cache.Cache
	Database *sql.DB
	ORM      *ent.Client
}

func NewContainer() *Container {
	c := new(Container)
	c.initWeb()
	c.initConfig()
	c.initCache()
	c.initDatabase()
	c.initORM()
	return c
}

func (c *Container) initWeb() {
	c.Web = echo.New()
}

func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		c.Web.Logger.Fatalf("failed to load configuration: %v", err)
	}
	c.Config = &cfg
}

func (c *Container) initCache() {
	cacheClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Config.Cache.Hostname, c.Config.Cache.Port),
		Password: c.Config.Cache.Password,
	})
	if _, err := cacheClient.Ping(context.Background()).Result(); err != nil {
		c.Web.Logger.Fatalf("failed to connect to cache server: %v", err)
	}
	cacheStore := store.NewRedis(cacheClient, nil)
	c.Cache = cache.New(cacheStore)
}

func (c *Container) initDatabase() {
	var err error

	addr := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		c.Config.Database.User,
		c.Config.Database.Password,
		c.Config.Database.Hostname,
		c.Config.Database.Database,
	)
	c.Database, err = sql.Open("pgx", addr)
	if err != nil {
		c.Web.Logger.Fatalf("failed to connect to database: %v", err)
	}
}

func (c *Container) initORM() {
	drv := entsql.OpenDB(dialect.Postgres, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		c.Web.Logger.Fatalf("failed to create database schema: %v", err)
	}
}
