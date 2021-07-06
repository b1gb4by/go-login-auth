package service

import (
	"api/util"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go/v4"
)

func GetUserID(t string, secret string) (string, error) {

	type Claims struct {
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		log.Println(err)
		return "", util.Errorf(util.ErrorCode10007, "", "%w", err)
	}

	claims := token.Claims.(*Claims)
	userID := claims.Issuer

	fmt.Println("userID: ", userID)

	return userID, nil
}
