package models

import "github.com/dgrijalva/jwt-go"

// Claims represents the JWT claims
type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}
