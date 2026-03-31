package main
import (
    "fmt"
    "github.com/Go_language/devops-healthcheck/checker"
    "github.com/Go_language/devops-healthcheck/modules"
)


func main(){
    fmt.Println("Welcome to Devops-healthcheck :")


    services := []models.Service{
        {Name: "gateway",  Port: 8080  , Healthy: true },
    	   {Name: "postgres",  Port: 5432 , Healthy: false },
    	   {Name: "frontend",  Port:  443 , Healthy:  true },
    }

    for _, svc  := range services {
	
		checker.PrintStatus(svc)
	}


}
