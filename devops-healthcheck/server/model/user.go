package model 

import (
	"time"
)

type User struct{
	ID int 
	Username string
	Email string
	Password string
	Created_at time.Time
}