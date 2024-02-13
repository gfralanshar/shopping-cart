package security

import (
	"fmt"
	"shopping-chart/api/v1/helper"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	helper.PanicIfError(err)

	if !token.Valid {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
