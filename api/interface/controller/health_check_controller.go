package controller

import (
	"api/infrastructure/database"
	"api/util"
	"net/http"

	"gorm.io/gorm"
)

type HealthCheckController struct {
	Auth *gorm.DB
}

func NewHealthCheckController(db database.Connection) *HealthCheckController {
	return &HealthCheckController{
		Auth: db.Auth,
	}
}

func (c *HealthCheckController) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	// DB 死活監視
	if err := c.DBHealthCheck(); err != nil {
		logger := util.NewStdLogger()
		logger.Fatalf("%s", err)
		responseJSON(w, http.StatusServiceUnavailable, "unable to connect to pf_encode")
	}

	responseJSON(w, http.StatusOK, "")
}

func (c *HealthCheckController) DBHealthCheck() error {
	sqlDB, err := c.Auth.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}
