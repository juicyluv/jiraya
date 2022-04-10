package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/logging"
	"sync"
)

// Config describes an application configuration structure.
type Config struct {
	// Http represents configuration for http server.
	Http struct {
		Port         string `yaml:"port" env-default:"8080"`
		ReadTimeout  int    `yaml:"readTimeout" env-default:"20"`
		WriteTimeout int    `yaml:"writeTimeout" env-default:"20"`
	} `yaml:"http"`
	GRPC struct {
		Port string `yaml:"port"`
	} `yaml:"grpc"`
	// DB represents configuration for database.
	DB struct {
		DSN               string `env:"DATABASE_DSN" env-required:"true"`
		RequestTimeout    int    `yaml:"requestTimeout" env-default:"5"`
		ConnectionTimeout int    `yaml:"connectionTimeout" env-default:"5"`
		ShutdownTimeout   int    `yaml:"shutdownTimeout" env-default:"5"`
	} `yaml:"database"`
	Logger struct {
		LogLevel string `yaml:"logLevel" env-default:"debug"`
	} `yaml:"logger"`
}

var cfg *Config
var once sync.Once

func New() (*Config, error) {
	logger := logging.Get()

	logger.Info("loading .env file")
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("could not load .env file: %v", err))
	}
	logger.Info("loaded .env file")

	once.Do(func() {
		logger.Info("reading application Config")

		cfg = &Config{}

		if err := cleanenv.ReadConfig("configs/dev.yml", cfg); err != nil {
			logger.Fatal(err.Error())
		}
	})
	logger.Info("done reading application Config")

	return cfg, nil
}

func Get() *Config {
	if cfg == nil {
		c, err := New()

		if err != nil {
			panic(err)
		}

		cfg = c
	}

	return cfg
}
