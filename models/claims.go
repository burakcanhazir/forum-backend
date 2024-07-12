package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims yapısı, JWT içinde taşınacak bilgileri içerir
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
