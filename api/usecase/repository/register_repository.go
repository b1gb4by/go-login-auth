package repository

import "api/domain/model"

type RegisterUserRepository interface {
	InsertData(user model.User) error
}
