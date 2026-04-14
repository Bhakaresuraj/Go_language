package main
import (
    "fmt"
    "time"
)

type Checker interface{
    Check()
} 


type HttpService struct{
    Name string
    URL string
    Healthy bool
}

func (h HttpService) Check(){
    // logic that checks  if service is healthy or not on the basis of h.URL

    fmt.Println("Checking if ",h.URL , "is HEalthy or not.....")
    time.Sleep(3*time.Second)
    h.Healthy =true

    fmt.Printf("HTTP => %s ,%s , %v \n",h.Name,h.URL,h.Healthy)

}

type 

func main(){
    fmt.Println("Welcome to Interfaces .......")

    service := HttpService {Name :"example.com",URL :"https://example.com",Healthy :false}
    services := []Checker{
             HttpService {Name :"example.com",URL :"https://example.com",Healthy :false},
    }

    for _,svc :=range services{
            
        svc.Check()
    }

    service.Check()

}
