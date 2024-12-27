package authentication

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var secret string = "ThisIsMySecretKey"

func GenerateToken(email string, ID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   ID,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	ok := parsedToken.Valid

	if !ok {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//
	if !ok {
		return 0, errors.New("couldn't get claims")
	}
	//
	//email := claims["email"].(string)
	ID := int64(claims["sub"].(float64))
	//expiresAt := claims["exp"].(string)

	return ID, nil
}
