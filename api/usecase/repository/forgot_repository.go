package repository

import "api/domain/model"

type ForgotRepository interface {
	InsertData(resetPassword model.ResetPassword) error
}
