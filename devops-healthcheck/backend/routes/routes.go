package Routes

import (
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/handlers"
	"net/http"
)

func RegisterRoutes(serviceHandler *handlers.ServiceHandler) {

	http.HandleFunc("/add", serviceHandler.AddService)
	http.HandleFunc("/update", serviceHandler.UpdateService)
	http.HandleFunc("/runall", serviceHandler.RunAll)
	http.HandleFunc("/delete", serviceHandler.DeleteHandler)
}
