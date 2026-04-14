package main
import (
    "fmt"
    "time"
)


type HttpService struct {
    Name string
    URL string
    Healthy bool
}

func Checker(h HttpService){
    h.Healthy =true

    fmt.Printf("HttpService :=>  %s  is checking (healthy or not) .\n",h.Name);
     time.Sleep(3* time.Second)
     fmt.Printf("HttpService :=>  %s  is Healthy : %v .\n",h.Name  ,h.Healthy);
}
func main(){
    fmt.Println("Hello go");

    service := HttpService{Name :"example.com" ,URL :"https://example.com",Healthy:false}

    Checker(service);
}
