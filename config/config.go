package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

type Env string

const (
	EnvLocal      Env = "local"
	EnvDevelop    Env = "dev"
	EnvStaging    Env = "staging"
	EnvQA         Env = "qa"
	EnvProduction Env = "prod"
)

// Config stores complete application configuration
type Config struct {
	HTTP  HTTPConfig
	App   AppConfig
	Cache CacheConfig
}

// HTTPConfig stores HTTP configuration
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`
	Port         uint16        `env:"HTTP_PORT,default=8000"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
}

// AppConfig stores application configuration
type AppConfig struct {
	Name          string `env:"APP_NAME,default=Goweb"`
	Environment   Env    `env:"APP_ENV,default=local"`
	EncryptionKey string `env:"APP_ENCRYPTION_KEY,default=?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf"`
}

type CacheConfig struct {
	Hostname string `env:"CACHE_HOSTNAME,default=localhost"`
	Port     uint16 `env:"CACHE_PORT,default=6379"`
	Password string `env:"CACHE_PASSWORD"`
	MaxAge   struct {
		StaticFile int `env:"CACHE_MAX_AGE_STATIC_FILE,default=15552000"`
		Page       int `env:"CACHE_STATIC_FILE_MAX_AGE,default=86400"`
	}
}

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
