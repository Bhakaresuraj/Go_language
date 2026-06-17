package Routes

import (
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/handlers"
	"net/http"
)

func RegisterRoutes(serviceHandler *handlers.ServiceHandler,
	authHandler *handlers.AuthHandler,
) { 
	// Authentication
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)
	//service routes
	http.HandleFunc("/add", serviceHandler.AddService)
	http.HandleFunc("/update", serviceHandler.UpdateService)
	http.HandleFunc("/runall", serviceHandler.RunAll)
	http.HandleFunc("/delete", serviceHandler.DeleteHandler)
}
