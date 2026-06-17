package utils

import (
	"os"
	"time"

	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	"github.com/golang-jwt/jwt/v5"
)
func GenerateToken(user model.User) (string, error) {

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
