package repository

import "api/domain/model"

type ResetRepository interface {
	SearchResetPassword(token string) (model.ResetPassword, error)
	UpdatePassword(email string, password []byte) error
}
