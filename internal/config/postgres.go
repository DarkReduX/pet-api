package config

import (
	"github.com/caarlos0/env"
)

type PostgreSQL struct {
	URL string `env:"POSTGRES_URL" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`
}

func NewPostgreSQLEnv() (*PostgreSQL, error) {
	cfg := &PostgreSQL{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
