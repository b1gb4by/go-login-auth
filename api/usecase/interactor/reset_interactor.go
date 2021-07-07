package interactor

import (
	"api/domain/model"
	"api/domain/service"
	"api/usecase/repository"
)

type resetInteractor struct {
	r repository.ResetRepository
}

type ResetInteractor interface {
	Reset(req model.ResetRequestParam) error
}

func NewResetInteractor(r repository.ResetRepository) ResetInteractor {
	return &resetInteractor{
		r: r,
	}
}

func (i *resetInteractor) Reset(req model.ResetRequestParam) error {

	if err := service.IsPasswordInputMatch(req.Password, req.ConfirmPassword); err != nil {
		return err
	}

	resetPassword, err := i.r.SearchResetPassword(req.Token)
	if err != nil {
		return err
	}

	password, err := service.GeneratedFromPassword(req.Password)
	if err != nil {
		return err
	}

	if err := i.r.UpdatePassword(resetPassword.Email, password); err != nil {
		return err
	}

	return nil
}
