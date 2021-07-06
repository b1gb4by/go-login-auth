package config

import (
	"api/util"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	User       string          `split_words:"true" required:"true"`
	Password   string          `split_words:"true" required:"true"`
	Host       string          `split_words:"true" required:"true"`
	Database   string          `split_words:"true" required:"true"`
	PORT       string          `required:"true" split_words:"true"`
	DBLogLevel logger.LogLevel `required:"true" split_words:"true"`
}

func NewDBConfig() *DBConfig {

	c := new(DBConfig)

	if err := envconfig.Process("mysql", c); err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}

	return c
}
