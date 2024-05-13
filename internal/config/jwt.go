package config

import (
	"github.com/caarlos0/env"
)

type JWT struct {
	ExpireMinutesAT int    `env:"JWT_EXPIRE_MINUTES_AT" envDefault:"60"`
	ExpireMinutesRT int    `env:"JWT_EXPIRE_MINUTES_RT" envDefault:"600"`
	SigningKeyAT    string `env:"JWT_SIGNING_KEY" envDefault:"AT_SIGNING_KEY"`
	SigningKeyRT    string `env:"JWT_SIGNING_KEY" envDefault:"RT_SIGNING_KEY"`
}

func NewJWT() (*JWT, error) {
	cfg := &JWT{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
