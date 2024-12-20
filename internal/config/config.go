package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
	"time"
)

type Config struct {
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Domain string `envconfig:"DOMAIN"`

	AuthTTL time.Duration `envconfig:"AUTH_TTL"`

	postgres.PGConfig
	redis.RedisConfig
}

func NewCfg() *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	if cfg.Host == "" {
		panic(errors.New("cfg is required"))
	}

	return &cfg
}
