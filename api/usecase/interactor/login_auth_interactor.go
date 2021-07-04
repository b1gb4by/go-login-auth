package interactor

import (
	"api/config"
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type loginAuthenticationInteractor struct {
	r  repository.LoginAuthenticationRepository
	jc *config.JWTConfig
}

type LoginAuthenticationInteractor interface {
	LoginAuthentication(req model.LoginAuthenticationRequestParam) (model.AcquisitionUser, string, error)
}

func NewLoginAuthenticationInteractor(
	r repository.LoginAuthenticationRepository,
	jc *config.JWTConfig,
) LoginAuthenticationInteractor {
	return &loginAuthenticationInteractor{
		r:  r,
		jc: jc,
	}
}

func (i *loginAuthenticationInteractor) LoginAuthentication(
	req model.LoginAuthenticationRequestParam,
) (model.AcquisitionUser, string, error) {
	var user model.AcquisitionUser
	var err error
	var token string

	user, err = i.r.SearchUser(req.Email)
	if err != nil {
		return user, token, err
	}

	if err := service.IsPasswordMatch(req.Password, user.Password); err != nil {
		return user, token, err
	}

	token, err = service.CreateJWT(user.ID, i.jc.Secret)
	if err != nil {
		return user, token, err
	}

	return user, token, nil
}
