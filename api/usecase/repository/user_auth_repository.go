package repository

import "api/domain/model"

type UserAuthenticationRepository interface {
	SearchUser(userID string) (model.AcquisitionUser, error)
}
