package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID      uint     `json:"user_id"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}
