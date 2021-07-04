package repository

import "api/domain/model"

type LoginAuthenticationRepository interface {
	SearchUser(email string) (model.AcquisitionUser, error)
}
