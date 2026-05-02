package main
import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
)

type Student struct {
    Name string   `json : "name"`
    Marks int      `json : "marks"`
    City string     `json :"city"`

}


func homeHandler(w http.ResponseWriter , r *http.Request){
    
    fmt.Println(r.Body)
    body ,err :=io.ReadAll(r.Body)
    if err!=nil{
        fmt.Println(err)
        return 
    }
    var stud Student
    err = json.Unmarshal(body ,&stud)
    if err != nil{
        fmt.Println("error in unMarshalling data")
    }
    
    fmt.Println(stud)
    fmt.Fprintln(w ,"Response from http server ...")

}

func main(){
    fmt.Println("Server us runnig on port :8080 ")
    http.HandleFunc("/",homeHandler)
    http.ListenAndServe(":8080",nil)

}
