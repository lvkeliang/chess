package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID uint `json:"ID"`
	jwt.StandardClaims
}
