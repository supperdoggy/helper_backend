package config

import (
	"github.com/caarlos0/env/v6"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type cfg struct {
	Port              int    `env:"PORT"`
	MongoUrl          string `env:"MONGO_URL"`
	EmailCheckService string `env:"EMAIL_CHECK_SERVICE"`
}

func NewConfig() (*cfg, error) {
	var c cfg
	err := env.Parse(&c)
	if err != nil {
		return nil, err
	}

	if err := ValidateConfig(c); err != nil {
		return nil, err
	}

	return &c, nil
}

func ValidateConfig(c cfg) error {
	return validation.ValidateStruct(&c, validation.Field(&c.Port, validation.Required),
		validation.Field(&c.MongoUrl, validation.Required), validation.Field(&c.EmailCheckService, validation.Required, is.URL))
}
