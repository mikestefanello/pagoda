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
	HTTP HTTPConfig
	App  AppConfig
}

// HTTPConfig stores HTTP configuration
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`
	Port         uint16        `env:"HTTP_PORT,default=8081"`
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

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
