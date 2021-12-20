package services

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
	"github.com/labstack/gommon/log"

	"goweb/config"
	"goweb/ent"
)

type Container struct {
	Web         *echo.Echo
	Config      *config.Config
	Cache       *cache.Cache
	cacheClient *redis.Client
	Database    *sql.DB
	ORM         *ent.Client
	Mail        *MailClient
	Auth        *AuthClient
	Templates   *TemplateRenderer
}

func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initORM()
	c.initMail()
	c.initAuth()
	c.initTemplateRenderer()
	return c
}

func (c *Container) Shutdown() error {
	if err := c.cacheClient.Close(); err != nil {
		return err
	}
	if err := c.ORM.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}

	return nil
}

func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

func (c *Container) initWeb() {
	c.Web = echo.New()

	// Configure logging
	switch c.Config.App.Environment {
	case config.EnvProduction:
		c.Web.Logger.SetLevel(log.WARN)
	default:
		c.Web.Logger.SetLevel(log.DEBUG)
	}
}

func (c *Container) initCache() {
	c.cacheClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Config.Cache.Hostname, c.Config.Cache.Port),
		Password: c.Config.Cache.Password,
	})
	if _, err := c.cacheClient.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Sprintf("failed to connect to cache server: %v", err))
	}
	cacheStore := store.NewRedis(c.cacheClient, nil)
	c.Cache = cache.New(cacheStore)
}

func (c *Container) initDatabase() {
	var err error

	getAddr := func(dbName string) string {
		return fmt.Sprintf("postgresql://%s:%s@%s/%s",
			c.Config.Database.User,
			c.Config.Database.Password,
			c.Config.Database.Hostname,
			dbName,
		)
	}

	c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.Database))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Check if this is a test environment
	if c.Config.App.Environment == config.EnvTest {
		// Drop the test database, ignoring errors in case it doesn't yet exist
		_, _ = c.Database.Exec("DROP DATABASE " + c.Config.Database.TestDatabase)

		// Create the test database
		if _, err = c.Database.Exec("CREATE DATABASE " + c.Config.Database.TestDatabase); err != nil {
			panic(fmt.Sprintf("failed to create test database: %v", err))
		}

		// Connect to the test database
		if err = c.Database.Close(); err != nil {
			panic(fmt.Sprintf("failed to close database connection: %v", err))
		}
		c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.TestDatabase))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}
	}
}

func (c *Container) initORM() {
	drv := entsql.OpenDB(dialect.Postgres, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed to create database schema: %v", err))
	}
}

func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config)
	if err != nil {
		panic(fmt.Sprintf("failed to create mail client: %v", err))
	}
}

func (c *Container) initAuth() {
	c.Auth = NewAuthClient(c.Config, c.ORM)
}

func (c *Container) initTemplateRenderer() {
	c.Templates = NewTemplateRenderer(c.Config)
}
