package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isSigningMethodHMAC := t.Method.(*jwt.SigningMethodHMAC); !isSigningMethodHMAC {
			return nil, errors.New("unknown signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok {
		return errors.New("invalid claim")
	}

	// email, userId := claims["email"].(string), claims["userId"].(int64)

	return nil

}
