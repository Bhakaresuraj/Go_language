package model

type ApiResponse struct {
	Success bool
	Message string
	Token  string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
