package main
import(
    "fmt"
)

func main(){
    fmt.Println("This is starting to go programming language")
    age :=20
    //age = "Suraj"
    fmt.Println(age)

    fmt.Println("Starting with functions ........")
    sum:=add(10,20)
    fmt.Println("Sum :",sum )
    division,err:=division(10,0)
    if err!=nil{
        fmt.Println("Error :",err)
        return
    }
    fmt.Println("Division :",division)
}

func add(a int ,b int) int{
    return a+b
}

func division(a float64,b float64)(float64 ,error){
    if b==0{
        return 0,fmt.Errorf("Cannot divide by 0 ....!")
    }
    return a/b ,nil
}



