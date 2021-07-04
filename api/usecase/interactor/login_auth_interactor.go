package interactor

import (
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type loginAuthenticationInteractor struct {
	r repository.LoginAuthenticationRepository
}

type LoginAuthenticationInteractor interface {
	LoginAuthentication(req model.LoginAuthenticationRequestParam) (string, error)
}

func NewLoginAuthenticationInteractor(r repository.LoginAuthenticationRepository) LoginAuthenticationInteractor {
	return &loginAuthenticationInteractor{
		r: r,
	}
}

func (i *loginAuthenticationInteractor) LoginAuthentication(req model.LoginAuthenticationRequestParam) (string, error) {
	var user model.AcquisitionUser
	var err error
	var token string

	user, err = i.r.SearchUser(req.Email)
	if err != nil {
		return token, err
	}

	if err := service.IsPasswordMatch(req.Password, user.Password); err != nil {
		return token, err
	}

	token, err = service.CreateJWT(user.ID)
	if err != nil {
		return token, err
	}

	return token, nil
}
