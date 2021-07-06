package interactor

import (
	"api/config"
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type userAuthenticationInteractor struct {
	r  repository.UserAuthenticationRepository
	jc *config.JWTConfig
}

type UserAuthenticationInteractor interface {
	UserAuthentication(token string) (model.AcquisitionUser, error)
}

func NewUserAuthenticationInteractor(
	r repository.UserAuthenticationRepository,
	jc *config.JWTConfig,
) UserAuthenticationInteractor {
	return &userAuthenticationInteractor{
		r:  r,
		jc: jc,
	}
}

func (i userAuthenticationInteractor) UserAuthentication(token string) (model.AcquisitionUser, error) {
	var user model.AcquisitionUser
	var userID string
	var err error

	userID, err = service.GetUserID(token, i.jc.Secret)
	if err != nil {
		return user, err
	}

	user, err = i.r.SearchUser(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
