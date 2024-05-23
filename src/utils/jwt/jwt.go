package jwt

import (
	"errors"
	"fmt"
	"projectsprintw4/src/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	UserId string
	Role   string
}

func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(config.JWT_EXP)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(v).Unix(),
		"iat":    time.Now().Unix(),
		"userId": payload.UserId,
		"role":   payload.Role,
	})

	token, err := t.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.JWT_SECRET), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return nil, errors.New("something went wrong.code 1a")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("something went wrong. code 1b")
	}

	return &TokenPayload{
		UserId: userId,
		Role:   role,
	}, nil
}
