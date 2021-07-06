package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type forgotGateway struct {
	db *gorm.DB
}

func NewForgotGateway(db database.Connection) repository.ForgotRepository {
	return &forgotGateway{
		db: db.Auth,
	}
}

func (g *forgotGateway) InsertData(resetPassword model.ResetPassword) error {
	const table = "reset_password"
	if err := g.db.Table(table).Create(&resetPassword).Error; err != nil {
		return util.Errorf(util.ErrorCode10002, "", "%w", err)
	}
	return nil
}
