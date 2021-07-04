package interactor

import (
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type registerUserInteractor struct {
	r repository.RegisterUserRepository
}

type RegisterUserInteractor interface {
	RegisterUser(req model.RegisterUserRequestParam) error
}

func NewRegisterUserInteractor(r repository.RegisterUserRepository) RegisterUserInteractor {
	return &registerUserInteractor{
		r: r,
	}
}

func (i *registerUserInteractor) RegisterUser(req model.RegisterUserRequestParam) error {

	var user model.User
	var err error

	user = model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if err = service.IsPasswordInputMatch(req.Password, req.ConfirmPassword); err != nil {
		return err
	}

	user.Password, err = service.GeneratedFromPassword(req.Password)
	if err != nil {
		return err
	}

	if err := i.r.InsertData(user); err != nil {
		return err
	}

	return nil
}
