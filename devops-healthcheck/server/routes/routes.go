package Routes

import (
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/handlers"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/middleware"
	"net/http"
)

func RegisterRoutes(serviceHandler *handlers.ServiceHandler,
	authHandler *handlers.AuthHandler,
) {
	// Authentication
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)
	//service routes
	http.HandleFunc("/add", middleware.AuthMiddleware(serviceHandler.AddService))
	http.HandleFunc("/update", middleware.AuthMiddleware(serviceHandler.UpdateService))
	http.HandleFunc("/services", middleware.AuthMiddleware(serviceHandler.GetAllServices))
	http.HandleFunc("/delete",  middleware.AuthMiddleware(serviceHandler.DeleteHandler))
}
