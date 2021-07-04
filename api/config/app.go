package config

import (
	"api/util"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port string `required:"true" split_words:"true"`
}

func NewAppConfig() *AppConfig {
	c := new(AppConfig)

	if err := envconfig.Process("api", c); err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}

	return c
}
