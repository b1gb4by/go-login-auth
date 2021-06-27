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
	Users *gorm.DB
}

func NewConnection(c *config.DBConfig) Connection {
	var dbConn Connection
	var err error

	schema := "%s:%s@tcp(mysql:%s)/%s?charset=utf8&parseTime=True"
	dsn := fmt.Sprintf(schema, c.User, c.Password, c.PORT, c.Database)

	dbConn.Users, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(c.DBLogLevel),
	})
	if err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}

	return dbConn
}

func (conn Connection) Close() {
	sqlDB, err := conn.Users.DB()
	if err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
	}
	if err = sqlDB.Close(); err != nil {
		return
	}
}
