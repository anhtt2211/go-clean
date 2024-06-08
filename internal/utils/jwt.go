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
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			// You can also add other custom claims here if needed
		},
	})

	// Set custom claims (id and username)
	token.Claims.(*UserClaims).ID = user.ID
	token.Claims.(*UserClaims).Username = user.Username

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
