package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"
	"errors"

	"gorm.io/gorm"
)

type loginAuthenticationGateway struct {
	db *gorm.DB
}

func NewLoginAuthenticationGateway(db database.Connection) repository.LoginAuthenticationRepository {
	return &loginAuthenticationGateway{
		db: db.Auth,
	}
}

func (g *loginAuthenticationGateway) SearchUser(email string) (model.AcquisitionUser, error) {
	const table = "user"
	var user model.AcquisitionUser

	if err := g.db.Table(table).Where("email = ?", email).First(&user).Error; err != nil {
		return user, util.Errorf(util.ErrorCode10003, "", "%w", err)
	}

	if user.ID == 0 {
		return user, util.Errorf(util.ErrorCode10004, "", "%w", errors.New("user not found."))
	}

	return user, nil
}
