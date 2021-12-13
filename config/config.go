package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

const (
	TemplateDir  = "views"
	TemplateExt  = ".gohtml"
	StaticDir    = "static"
	StaticPrefix = "files"
)

type Env string

const (
	EnvLocal      Env = "local"
	EnvTest       Env = "test"
	EnvDevelop    Env = "dev"
	EnvStaging    Env = "staging"
	EnvQA         Env = "qa"
	EnvProduction Env = "prod"
)

type (
	// Config stores complete application configuration
	Config struct {
		HTTP     HTTPConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Hostname     string        `env:"HTTP_HOSTNAME"`
		Port         uint16        `env:"HTTP_PORT,default=8000"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
		IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
	}

	// AppConfig stores application configuration
	AppConfig struct {
		Name          string        `env:"APP_NAME,default=Goweb"`
		Environment   Env           `env:"APP_ENV,default=local"`
		EncryptionKey string        `env:"APP_ENCRYPTION_KEY,default=?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf"`
		Timeout       time.Duration `env:"APP_TIMEOUT,default=20s"`
	}

	CacheConfig struct {
		Hostname   string `env:"CACHE_HOSTNAME,default=localhost"`
		Port       uint16 `env:"CACHE_PORT,default=6379"`
		Password   string `env:"CACHE_PASSWORD"`
		Expiration struct {
			StaticFile time.Duration `env:"CACHE_EXPIRATION_STATIC_FILE,default=4380h"`
			Page       time.Duration `env:"CACHE_EXPIRATION_PAGE,default=24h"`
		}
	}

	DatabaseConfig struct {
		Hostname     string `env:"DB_HOSTNAME,default=localhost"`
		Port         uint16 `env:"DB_PORT,default=5432"`
		User         string `env:"DB_USER,default=admin"`
		Password     string `env:"DB_PASSWORD,default=admin"`
		Database     string `env:"DB_NAME,default=app"`
		TestDatabase string `env:"DB_NAME_TEST,default=app_test"`
	}
)

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
