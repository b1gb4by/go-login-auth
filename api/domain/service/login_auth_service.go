package service

import (
	"api/util"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

func IsPasswordMatch(inputPassword string, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(password, []byte(inputPassword)); err != nil {
		return util.Errorf(util.ErrorCode10005, "", "%w", err)
	}
	return nil
}

func CreateJWT(id int, secret string) (string, error) {
	expiresAt := jwt.NewTime(float64(time.Now().Add(time.Hour * 24).Unix()))
	mySigningKey := []byte(secret)

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
