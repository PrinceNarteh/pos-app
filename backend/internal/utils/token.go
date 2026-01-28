package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string
}

func generateToken(secret, expTime string, user *models.User) (string, error) {
	duration, err := time.ParseDuration(expTime)
	if err != nil {
		return "", fmt.Errorf("invalid token expiration time: %w", err)
	}

	now := time.Now()
	exp := now.Add(duration)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(user.ID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signed, nil
}

func GenerateAccessToken(user *models.User) (string, error) {
	return generateToken(config.Env.Jwt.AccessSecret, config.Env.Jwt.AccessExpirationTime, user)
}

func GenerateRefreshToken(user *models.User) (string, error) {
	return generateToken(config.Env.Jwt.RefreshSecret, config.Env.Jwt.RefreshExpirationTime, user)
}

func ParseToken(tokenStr string, isAccessToken bool) (*Claims, error) {
	claims := new(Claims)
	secret := config.Env.Jwt.RefreshSecret
	if isAccessToken {
		secret = config.Env.Jwt.AccessSecret
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
