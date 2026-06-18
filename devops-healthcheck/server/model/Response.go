package model

type ApiResponse struct {
	Success bool
	Message string
	Token  string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}