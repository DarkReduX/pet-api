package config

import (
	"github.com/caarlos0/env"
)

type Postgres struct {
	URL string `env:"POSTGRES_URL" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`
}

func NewPostgres() (*Postgres, error) {
	cfg := &Postgres{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
