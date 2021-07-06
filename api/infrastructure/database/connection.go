package database

import (
	"api/config"
	"api/util"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	Auth *gorm.DB
}

func NewConnection(c *config.DBConfig) Connection {
	var dbConn Connection
	var err error

	schema := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	dsn := fmt.Sprintf(schema, c.User, c.Password, c.Host, c.PORT, c.Database)

	dbConn.Auth, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(c.DBLogLevel),
	})
	if err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}

	return dbConn
}

func (conn Connection) Close() {
	sqlDB, err := conn.Auth.DB()
	if err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}
	if err = sqlDB.Close(); err != nil {
		return
	}
}
