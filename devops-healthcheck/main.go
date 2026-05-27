package main

import (
	"encoding/json"
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/model"
	"github.com/rs/cors"
	"io"
	"log"
	"net/http"
	"time"
)
var db *database.Store
func AddService(w http.ResponseWriter, r *http.Request) {
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
	err = db.Save(service)
	if err != nil {
		http.Error(w, "Error Saving Service ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
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
	err = db.UpdateService(service)
	if err != nil {
		http.Error(w, "Error Updating Services : ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func RunAll(w http.ResponseWriter, r *http.Request) {

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
	services, err := db.GetAllServices(service.UserID)
	if err != nil {
		http.Error(w, "Error while fetching for update :", http.StatusInternalServerError)
		return
	}
	fmt.Println("Run all Called")
	for _, ser := range services {
		avg_time := time.Now()
		ser.Healthy, ser.StatusCode = ser.CheckHealth()
		responseTime := time.Since(avg_time)
		ser.Checked_at = time.Now()
		ser.Response_time = responseTime.Milliseconds()
		err = db.UpdateServiceStatus(ser)
		if err != nil {
			fmt.Println("Update Error :", err)
			continue
		}
	}
	services, err = db.GetAllServices(service.UserID)
	if err != nil {
		http.Error(w, "Error Geting Services : ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(services)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
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
	err = db.DeleteService(service.ID)

	if err != nil {
		http.Error(w, "Service not deleted ", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Successfully deleted Service")
}

func main() {
	fmt.Println("Server is Running on Port :8080")
	dbUrl := "postgres://suraj:Bhakare@localhost:5432/mydb"
	db = database.NewStore(dbUrl)
	fmt.Println("Database Connected Successfully")
	err := db.RunMigration()
	if err != nil {
		log.Fatal("Unable to run migration :", err)
	}

	// Routes

	http.HandleFunc("/add", AddService)
	http.HandleFunc("/update", UpdateService)
	http.HandleFunc("/runall", RunAll)
	http.HandleFunc("/delete", DeleteHandler)

	// CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Wrap Default ServeMux
	handler := c.Handler(http.DefaultServeMux)
	// Start Server

	log.Fatal(http.ListenAndServe(":8080", handler))
}
