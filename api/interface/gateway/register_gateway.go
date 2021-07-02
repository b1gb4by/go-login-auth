package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type registerUserGateway struct {
	db    *gorm.DB
	table string
}

func NewRegisterUserGateway(db database.Connection, tn string) repository.RegisterUserRepository {
	return &registerUserGateway{
		db:    db.Users,
		table: tn,
	}
}

func (g *registerUserGateway) InsertData(user model.User) error {
	if err := g.db.Table(g.table).Create(&user).Error; err != nil {
		return util.Errorf(util.ErrorCode10002, "", "%w", err)
	}
	return nil
}
