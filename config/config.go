package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type environment string

const (
	// EnvLocal represents the local environment.
	EnvLocal environment = "local"

	// EnvTest represents the test environment.
	EnvTest environment = "test"

	// EnvDevelopment represents the development environment.
	EnvDevelopment environment = "dev"

	// EnvStaging represents the staging environment.
	EnvStaging environment = "staging"

	// EnvQA represents the qa environment.
	EnvQA environment = "qa"

	// EnvProduction represents the production environment.
	EnvProduction environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env environment) {
	if err := os.Setenv("PAGODA_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration.
	Config struct {
		HTTP     HTTPConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
		Files    FilesConfig
		Tasks    TasksConfig
		Mail     MailConfig
	}

	// HTTPConfig stores HTTP configuration.
	HTTPConfig struct {
		Hostname        string
		Port            uint16
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		IdleTimeout     time.Duration
		ShutdownTimeout time.Duration
		TLS             struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	// AppConfig stores application configuration.
	AppConfig struct {
		Name          string
		Host          string
		Environment   environment
		EncryptionKey string
		Timeout       time.Duration
		PasswordToken struct {
			Expiration time.Duration
			Length     int
		}
		EmailVerificationTokenExpiration time.Duration
	}

	// CacheConfig stores the cache configuration.
	CacheConfig struct {
		Capacity   int
		Expiration struct {
			PublicFile time.Duration
		}
	}

	// DatabaseConfig stores the database configuration.
	DatabaseConfig struct {
		Driver         string
		Connection     string
		TestConnection string
	}

	// FilesConfig stores the file system configuration.
	FilesConfig struct {
		Directory string
	}

	// TasksConfig stores the tasks configuration.
	TasksConfig struct {
		Goroutines      int
		ReleaseAfter    time.Duration
		CleanupInterval time.Duration
		ShutdownTimeout time.Duration
	}

	// MailConfig stores the mail configuration.
	MailConfig struct {
		Hostname    string
		Port        uint16
		User        string
		Password    string
		FromAddress string
	}
)

// GetConfig loads and returns configuration.
func GetConfig() (Config, error) {
	var c Config

	// Load the config file.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// Load env variables.
	viper.SetEnvPrefix("pagoda")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}
