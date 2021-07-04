package service

import (
	"api/util"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func IsPasswordMatch(p string, cp string) error {
	if p != cp {
		return util.Errorf(util.ErrorCode10000, "", "%w", errors.New("password and retype password do not match."))
	}
	return nil
}

func GeneratedFromPassword(p string) ([]byte, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	if err != nil {
		return nil, util.Errorf(util.ErrorCode10001, "", "%w", err)
	}
	return password, nil
}
