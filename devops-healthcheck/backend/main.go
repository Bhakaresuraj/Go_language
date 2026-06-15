package main

import (
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/handlers"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/middleware"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/routes"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/worker"
	"log"
	"net/http"

)
func main() {
	fmt.Println("Server is Running on Port :8080")
	dbUrl := "postgres://suraj:Bhakare@localhost:5432/mydb"
	db := database.NewStore(dbUrl)
	fmt.Println("Database Connected Successfully")
	err := db.RunMigration()
	if err != nil {
		log.Fatal("Unable to run migration :", err)
	}
	serviceHandler := &handlers.ServiceHandler{
		DB: db,
	}
	go worker.StartBackgroundWorker(db);
	Routes.RegisterRoutes(serviceHandler)
	c := middleware.SetupCors()
	handler := c.Handler(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
