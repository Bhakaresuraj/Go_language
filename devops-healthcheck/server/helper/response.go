package helper

import (
	"encoding/json"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, response model.ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
