package config

type JWT struct {
	ExpireMinutesAT int    `env:"JWT_EXPIRE_MINUTES_AT" envDefault:"60"`
	ExpireMinutesRT int    `env:"JWT_EXPIRE_MINUTES_RT" envDefault:"600"`
	SigningKeyAT    string `env:"JWT_SIGNING_KEY" envDefault:"AT_SIGNING_KEY"`
	SigningKeyRT    string `env:"JWT_SIGNING_KEY" envDefault:"RT_SIGNING_KEY"`
}
