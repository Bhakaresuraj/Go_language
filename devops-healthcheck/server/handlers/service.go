package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"time"

	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/helper"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
)

type ServiceHandler struct {
	DB *database.Store
}

func (h *ServiceHandler) AddService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed",
		}
		helper.SendResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error reading body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	var service model.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error unmarshalling body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	userID := r.Context().Value("user_id").(int)
	fmt.Println("Logged User:", userID)
	avg_time := time.Now()
	service.UserID = userID
	service.Healthy, service.StatusCode = service.CheckHealth()
	responseTime := time.Since(avg_time)
	service.Response_time = responseTime.Milliseconds()
	service.Checked_at = time.Now()
	err = h.DB.Save(service)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Saving Service ",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return

	}
	response := model.ApiResponse{
		Success: true,
		Message: "Service Saved Successfully...!",
	}
	helper.SendResponse(w, http.StatusOK, response)
}

func (h *ServiceHandler) UpdateService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed",
		}
		helper.SendResponse(w, http.StatusMethodNotAllowed, response)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error reading body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	var service model.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error unmarshalling body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	userID := r.Context().Value("user_id").(int)
	avg_time := time.Now()
	service.UserID = userID
	service.Healthy, service.StatusCode = service.CheckHealth()
	responseTime := time.Since(avg_time)
	service.Checked_at = time.Now()
	service.Response_time = responseTime.Milliseconds()
	fmt.Println(service)
	err = h.DB.UpdateService(service)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Updating Service ",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	response := model.ApiResponse{
		Success: true,
		Message: "Service Updated Successfully...!",
	}

	helper.SendResponse(w, http.StatusOK, response)
}

func (h *ServiceHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed",
		}
		helper.SendResponse(w, http.StatusMethodNotAllowed, response)
		return
	}
	userID := r.Context().Value("user_id").(int)
	fmt.Println("Request from browser for services")
	services, err := h.DB.GetAllServices(userID)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Geting Services : ",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	response := model.ApiResponse{
		Success: true,
		Message: "Successfully Fetched Data !",
		Data:    services,
	}
	helper.SendResponse(w, http.StatusOK, response)
}

func (h *ServiceHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed",
		}
		helper.SendResponse(w, http.StatusMethodNotAllowed, response)
		return
	}
	// userID := r.Context().Value("user_id").(int)
	var service model.Service
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error reading body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	err = json.Unmarshal(body, &service)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error unmarshalling body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	// fmt.Println(service,userID)
	err = h.DB.DeleteService(service.ID)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Deleting Service ",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	response := model.ApiResponse{
		Success: true,
		Message: "Service Deleted Successfully...!",
	}

	helper.SendResponse(w, http.StatusOK, response)
}
