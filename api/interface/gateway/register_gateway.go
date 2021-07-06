package gateway

import (
	"api/domain/model"
	"api/infrastructure/database"
	"api/usecase/repository"
	"api/util"

	"gorm.io/gorm"
)

type registerUserGateway struct {
	db *gorm.DB
}

func NewRegisterUserGateway(db database.Connection) repository.RegisterUserRepository {
	return &registerUserGateway{
		db: db.Auth,
	}
}

func (g *registerUserGateway) InsertData(user model.User) error {
	const table = "user"
	if err := g.db.Table(table).Create(&user).Error; err != nil {
		return util.Errorf(util.ErrorCode10002, "", "%w", err)
	}
	return nil
}
