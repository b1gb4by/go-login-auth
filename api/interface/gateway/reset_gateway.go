package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type resetGateway struct {
	db *gorm.DB
}

func NewResetGateway(db database.Connection) repository.ResetRepository {
	return &resetGateway{
		db: db.Auth,
	}
}

func (g *resetGateway) SearchResetPassword(token string) (model.ResetPassword, error) {
	const table = "reset_password"
	var resetPassword model.ResetPassword

	if err := g.db.Table(table).Where("token = ?", token).Last(&resetPassword).Error; err != nil {
		return resetPassword, util.Errorf(util.ErrorCode10009, "", "%w", err)
	}

	return resetPassword, nil
}

func (g *resetGateway) UpdatePassword(email string, password []byte) error {
	const table = "users"

	if err := g.db.Table(table).Where("email = ?", email).Update("password", password).Error; err != nil {
		return util.Errorf(util.ErrorCode10010, "", "%w", err)
	}

	return nil
}
