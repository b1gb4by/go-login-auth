package interactor

import (
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type resetPasswordInteractor struct {
	r repository.ResetPasswordRepository
}

type ResetPasswordInteractor interface {
	ResetPassword(req model.ResetPasswordRequestParam) error
}

func NewResetPasswordInteractor(r repository.ResetPasswordRepository) ResetPasswordInteractor {
	return &resetPasswordInteractor{
		r: r,
	}
}

func (i *resetPasswordInteractor) ResetPassword(req model.ResetPasswordRequestParam) error {

	var resetPassword = model.ResetPassword{
		Email: req.Email,
		Token: service.CreateRandomString(),
	}

	if err := i.r.InsertData(resetPassword); err != nil {
		return err
	}

	return nil
}
