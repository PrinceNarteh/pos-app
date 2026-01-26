package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string
}

func GenerateToken(user *models.User) (string, error) {
	now := time.Now()
	exp := now.Add(time.Hour)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(user.ID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte("hello"))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signed, nil
}
