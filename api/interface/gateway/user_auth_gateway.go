package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type userAuthenticationGateway struct {
	db *gorm.DB
}

func NewUserAuthenticationGateway(db database.Connection) repository.UserAuthenticationRepository {
	return &userAuthenticationGateway{
		db: db.Auth,
	}
}

func (g *userAuthenticationGateway) SearchUser(userID string) (model.AcquisitionUser, error) {
	const table = "user"
	var user model.AcquisitionUser

	if err := g.db.Table(table).Where("id = ?", userID).First(&user).Error; err != nil {
		return user, util.Errorf(util.ErrorCode10003, "", "%w", err)
	}

	return user, nil
}
