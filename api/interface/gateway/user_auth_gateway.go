package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type userAuthenticationGateway struct {
	db    *gorm.DB
	table string
}

func NewUserAuthenticationGateway(db database.Connection, tn string) repository.UserAuthenticationRepository {
	return &userAuthenticationGateway{
		db:    db.Users,
		table: tn,
	}
}

func (g *userAuthenticationGateway) SearchUser(userID string) (model.AcquisitionUser, error) {
	var user model.AcquisitionUser

	if err := g.db.Table(g.table).Where("id = ?", userID).First(&user).Error; err != nil {
		return user, util.Errorf(util.ErrorCode10003, "", "%w", err)
	}

	return user, nil
}
