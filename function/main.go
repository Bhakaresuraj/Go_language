package main
import(
    "fmt"
)


func checkPort(port int )(string, error){
    if port <=0 || port >=65535{
        return "",fmt.Errorf("This is invalid port : %d",port)
    } else {
        return fmt.Sprintf("port %d is valid port .",port ) ,nil
    }
}
func main(){
    fmt.Println("Learning Functions in Go language ..:")

    //create and assign varaible

    validPort :=8080
    msg ,err:=checkPort(validPort)
    if err != nil{
        fmt.Println("Error :",err)
        return 
    }

    fmt.Println(msg) 

    //create and assign varaible

    inValidPort :=-10
    wmsg ,werr:=checkPort(inValidPort)
    if werr != nil{
        fmt.Println("Error :",werr)
        return 
    }

    fmt.Println(wmsg) 

    

}
