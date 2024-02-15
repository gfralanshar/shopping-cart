package security

import (
	"os"
	"shopping-chart/api/v1/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("API_KEY"))

func CreateToken(customer_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   customer_id,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	helper.PanicIfError(err)
	return tokenString, nil
}
