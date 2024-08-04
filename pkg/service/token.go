package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/models"
)

func (a *AdminService) GenerateToken(userName string) (string, error) {
	admin, err := a.AdminRepo.FindAdmin(userName)
	if err != nil {
		return "", err
	}

	claims := &models.Claims{
		AdminID: admin.AdminID,
		Email:   admin.Email,
		Role:    admin.Role,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("q3e67yajhsdb4")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
