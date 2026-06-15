package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "sync"
	"time"

	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

type ServiceHandler struct {
	DB *database.Store
}

func (h *ServiceHandler) AddService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	var service model.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		http.Error(w, "Error unmarshalling body", http.StatusInternalServerError)
		return
	}

	avg_time := time.Now()
	service.Healthy, service.StatusCode = service.CheckHealth()
	responseTime := time.Since(avg_time)
	service.Response_time = responseTime.Milliseconds()
	service.Checked_at = time.Now()
	err = h.DB.Save(service)
	if err != nil {
		http.Error(w, "Error Saving Service ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	var service model.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		http.Error(w, "Error Unmarshalling Body :", http.StatusInternalServerError)
		return
	}
	avg_time := time.Now()
	service.Healthy, service.StatusCode = service.CheckHealth()
	responseTime := time.Since(avg_time)
	service.Checked_at = time.Now()
	service.Response_time = responseTime.Milliseconds()
	err = h.DB.UpdateService(service)
	if err != nil {
		http.Error(w, "Error Updating Services : ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func (h *ServiceHandler) RunAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in reading Body", http.StatusInternalServerError)
		return
	}
	var service model.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		http.Error(w, "Error while Unmarshalling body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Request from browser for services")
	services, err := h.DB.GetAllServices(service.UserID)
	if err != nil {
		http.Error(w, "Error Geting Services : ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(services)
}

func (h *ServiceHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed !", http.StatusMethodNotAllowed)
		return
	}
	var service model.Service
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in Reading Body ", http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &service)
	if err != nil {
		http.Error(w, "Error in Unmarshalling Body ", http.StatusInternalServerError)

	}
	err = h.DB.DeleteService(service.ID)

	if err != nil {
		http.Error(w, "Service not deleted ", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Successfully deleted Service")
}
