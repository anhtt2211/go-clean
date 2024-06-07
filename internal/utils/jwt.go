package utils

import (
	"My-Clean/internal/domain/entities"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your-secret-key") // Replace with your secret key

type UserClaims struct {
	entities.User
	jwt.StandardClaims
}

func (u UserClaims) Valid() error {
	if u.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}
	return nil
}

func GenerateJWT(user entities.User) (string, error) {
	return "", nil
}

func ValidateJWT(tokenStr string) (entities.User, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return entities.User{}, err
	}

	if !token.Valid {
		return entities.User{}, errors.New("invalid token")
	}

	return claims.User, nil
}
