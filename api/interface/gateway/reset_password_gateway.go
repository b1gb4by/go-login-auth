package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type resetPasswordGateway struct {
	db *gorm.DB
}

func NewResetPasswordGateway(db database.Connection) repository.ResetPasswordRepository {
	return &resetPasswordGateway{
		db: db.Auth,
	}
}

func (g *resetPasswordGateway) InsertData(resetPassword model.ResetPassword) error {
	const table = "reset_password"
	if err := g.db.Table(table).Create(&resetPassword).Error; err != nil {
		return util.Errorf(util.ErrorCode10002, "", "%w", err)
	}
	return nil
}
