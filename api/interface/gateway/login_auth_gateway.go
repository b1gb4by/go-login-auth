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
	db    *gorm.DB
	table string
}

func NewLoginAuthenticationGateway(db database.Connection, tn string) repository.LoginAuthenticationRepository {
	return &loginAuthenticationGateway{
		db:    db.Users,
		table: tn,
	}
}

func (g *loginAuthenticationGateway) SearchUser(email string) (model.AcquisitionUser, error) {
	var user model.AcquisitionUser

	if err := g.db.Table(g.table).Where("email = ?", email).First(&user).Error; err != nil {
		return user, util.Errorf(util.ErrorCode10003, "", "%w", err)
	}

	if user.ID == 0 {
		return user, util.Errorf(util.ErrorCode10004, "", "%w", errors.New("user not found."))
	}

	return user, nil
}
