package database

import (
	"fmt"

	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
)

func (s *Store) SaveUser(user model.User) error {
	query := `INSERT INTO users (username,email,password) VALUES($1,$2,$3)`
	_, err := s.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByEmail(email string) (bool, error ,model.User){
	query := `SELECT id, username, email, password, created_at FROM users WHERE email=$1`
	var user model.User
	err := s.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_at)
	if err != nil {
		return false, err ,user
	}
	fmt.Println("Get user by email :", user)
	return true, nil, user
}
