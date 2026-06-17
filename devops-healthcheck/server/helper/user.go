package helper

import (
	"errors"
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUserRequest(user model.User) bool {

	if user.Email == "" || user.Password == "" {
		return false
	}
	return true
}

func AuthenticateUser(password string, hashed_password string) bool {

	
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			fmt.Println("Login failed: Incorrect password.")
			return false
		}
	}

	return true
}
