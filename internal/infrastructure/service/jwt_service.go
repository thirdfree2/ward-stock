package service

import (
	"fmt"
	"os"
	"time"

	"ward-stock-backend/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userID uint, email string, perms []string) (string, error)
}

type jwtServiceImpl struct {
	secretKey []byte
}

func NewJwtService() JwtService {
	secret := os.Getenv("JWT_SECRET")
	fmt.Println(secret)
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return &jwtServiceImpl{
		secretKey: []byte(secret),
	}
}

func (j *jwtServiceImpl) GenerateToken(userID uint, email string, perms []string) (string, error) {
	claims := auth.CustomClaims{
		UserID:      userID,
		Email:       email,
		Permissions: perms,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}
