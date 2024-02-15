package security

import (
	"fmt"
	"net/http"
	"shopping-chart/api/v1/helper"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func VerifyToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	// You can use `token` if needed for further processing

	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenId(r *http.Request) (int, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
		helper.PanicIfError(err)
		return int(id), nil
	}

	return 0, nil
}
