package config

import (
	"api/util"

	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	Secret string `split_words:"true" required:"true"`
}

func NewJWTConfig() *JWTConfig {

	c := new(JWTConfig)

	if err := envconfig.Process("jwt", c); err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}

	return c
}
