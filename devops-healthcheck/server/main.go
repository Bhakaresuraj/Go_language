package main

import (
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/handlers"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/middleware"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/routes"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/worker"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server is Running on Port :8080")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	db := database.NewStore(dbUrl)
	fmt.Println("Database Connected Successfully")
	err = db.RunMigration()
	if err != nil {
		log.Fatal("Unable to run migration :", err)
	}


	serviceHandler := &handlers.ServiceHandler{
		DB: db,
	}
	authHandler := &handlers.AuthHandler{
		DB: db,
	}

	
	// Starting background worker
	go worker.StartBackgroundWorker(db)

	// Routes creating for services and authentication
	Routes.RegisterRoutes(serviceHandler, authHandler)

	// middleware for cors
	c := middleware.SetupCors()
	handler := c.Handler(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handler))
}
