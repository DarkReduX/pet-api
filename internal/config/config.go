package config

import "github.com/caarlos0/env/v11"

type Config struct {
	ProductServiceEndpoint string `env:"PRODUCT_SERVICE_ENDPOINT" envDefault:"http://localhost:8081"`

	JWT
	Postgres
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
