package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io"
	"time"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/model"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/database"
)
var db *database.Store

func HealthCheckHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Welcome to DevOps Healthcheck.....")
	if r.Method != http.MethodPost{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	body,err:=io.ReadAll(r.Body)
	if err != nil{
		http.Error(w,"Error reading body",http.StatusInternalServerError)
		return
	}
	var service model.Service
	err = json.Unmarshal(body,&service)
	if err != nil{
		http.Error(w,"Error unmarshalling body",http.StatusInternalServerError)
		return
	}
	service.Healthy,service.StatusCode = service.CheckHealth()
	service.Checked_at=time.Now()
	err=db.Save(service)
	if err !=nil{
		http.Error(w,"Error Saving Service ",http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(service)
}
func GetAllServices(w http.ResponseWriter ,r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",http.StatusInternalServerError)
		return
	}
	body,err:=io.ReadAll(r.Body)
	var service model.Service
	 err =json.Unmarshal(body,&service)
	if err != nil{
		http.Error(w,"Error Unmarshalling Body :",http.StatusInternalServerError)
		return 
	}
	services,err:= db.GetAllServices(service.UserID)	
	if err!=nil{
		http.Error(w,"Error Geting Services : ",http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(services)
}

func main(){
	fmt.Println("Server is Running on Port :8080")
	dbUrl := "postgres://suraj:Bhakare@localhost:5432/mydb"
	db =database.NewStore(dbUrl)
	fmt.Println("Database Connected Successfully")
	err:=db.RunMigration()
	if err != nil{
		log.Fatal("Unable to run migration :",err)
	}
	fmt.Println("Migration Run Successfully and Table Created...!")
	http.HandleFunc("/check",HealthCheckHandler)
	http.HandleFunc("/services",GetAllServices)
	http.ListenAndServe(":8080", nil)
}
