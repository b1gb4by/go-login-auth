package repository

import "api/domain/model"

type ResetPasswordRepository interface {
	InsertData(resetPassword model.ResetPassword) error
}
