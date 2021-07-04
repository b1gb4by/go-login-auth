package service

import (
	"api/util"
	"errors"
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

func IsPasswordMatch(inputPassword string, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(password, []byte(inputPassword)); err != nil {
		return util.Errorf(util.ErrorCode10005, "", "%w", errors.New("password and retype password do not match."))
	}
	return nil
}

func CreateJWT(id int) (string, error) {

	var expiresAt = jwt.NewTime(60 * 60 * 24)
	var mySigningKey = []byte("secret")

	claims := &jwt.StandardClaims{
		Issuer:    strconv.Itoa(id),
		ExpiresAt: expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", util.Errorf(util.ErrorCode10006, "", "%w", err)
	}

	return ss, nil
}
