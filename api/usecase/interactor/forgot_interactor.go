package interactor

import (
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type forgotInteractor struct {
	r repository.ForgotRepository
}

type ForgotInteractor interface {
	Forgot(req model.ForgotRequestParam) error
}

func NewForgotInteractor(r repository.ForgotRepository) ForgotInteractor {
	return &forgotInteractor{
		r: r,
	}
}

func (i *forgotInteractor) Forgot(req model.ForgotRequestParam) error {

	token := service.CreateRandomString()

	var resetPassword = model.ResetPassword{
		Email: req.Email,
		Token: token,
	}

	if err := i.r.InsertData(resetPassword); err != nil {
		return err
	}

	if err := service.SendToSMTP(req.Email, token); err != nil {
		return err
	}

	return nil
}
