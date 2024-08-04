package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	AdminID uint   `json:"userid"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Claims  jwt.StandardClaims
}

// Valid implements jwt.Claims.
func (*Claims) Valid() error {
	panic("unimplemented")
}
