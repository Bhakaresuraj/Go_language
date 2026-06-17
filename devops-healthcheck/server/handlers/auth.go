package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/helper"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/utils"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type AuthHandler struct {
	DB *database.Store
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed ",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error reading User body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return

	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Unmarshalling User body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	if !helper.ValidateUserRequest(user) {
		response := model.ApiResponse{
			Success: false,
			Message: "Enter valid Information(email,Password)..!",
		}
		helper.SendResponse(w, http.StatusBadRequest, response)
		return
	}
	isExists, err, _ := a.DB.GetUserByEmail(user.Email)
	if isExists == true {
		response := model.ApiResponse{
			Success: false,
			Message: "Email Already exists ,Please Try with new Email ..!",
		}
		helper.SendResponse(w, http.StatusBadRequest, response)
		return
	}
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error during Hashing Passoword")
		response := model.ApiResponse{
			Success: false,
			Message: "Server Error",
		}
		helper.SendResponse(w, http.StatusInternalServerError,response)
		return
	}
	user.Password = string(hashpassword)
	fmt.Println(user)
	err = a.DB.SaveUser(user)
	if err != nil {
		fmt.Println("Error during Saving User in database")
		response := model.ApiResponse{
			Success: false,
			Message: "Server Error",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return
	}
	response := model.ApiResponse{
			Success: true,
			Message: "User Registered Successfully..!",
		}
	helper.SendResponse(w, http.StatusOK,response)

}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := model.ApiResponse{
			Success: false,
			Message: "Method not allowed ",
		}
		helper.SendResponse(w, http.StatusInternalServerError,response)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error reading User body",
		}
		helper.SendResponse(w, http.StatusInternalServerError, response)
		return

	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Error Unmarshalling User body",
		}
		helper.SendResponse(w, http.StatusInternalServerError,response)
		return
	}
	if !helper.ValidateUserRequest(user) {
		response := model.ApiResponse{
			Success: false,
			Message: "Enter valid Information(email,Password)..!",
		}
		helper.SendResponse(w, http.StatusBadRequest, response)
		return
	}
	isExists, err, present_user := a.DB.GetUserByEmail(user.Email)
	if isExists == false {
		response := model.ApiResponse{
			Success: false,
			Message: "User Not Found . Please Sign Up",
		}
		helper.SendResponse(w, http.StatusBadRequest, response)
		return
	}
	isLogedIN := helper.AuthenticateUser(user.Password, present_user.Password)
	if !isLogedIN {
		response := model.ApiResponse{
			Success: false,
			Message:"Invalid Credentials",
		}
		helper.SendResponse(w, http.StatusBadRequest,response)
		return
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Invalid Credentials",
		}
		helper.SendResponse(w, http.StatusBadRequest, response)
		return
	}
	response := model.ApiResponse{
			Success: true,
			Message: "Login Successfully..!",
			Token: token,
		}
	helper.SendResponse(w, http.StatusOK, response)

}
