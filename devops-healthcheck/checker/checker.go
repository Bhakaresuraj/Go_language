func package checker
import (

  "fmt"
  "github.com/fatih/color"
  "github.com/Go_language/devops-healthcheck/models"
)

func PrintStatus(s models.Service){
    status := "Healthy"
    if s.Healthy ==false{
        status ="UnHealthy"
        msg :=fmt.Sprintf("Name:%s  | Port :%d | %s ",s.Name,s.Port,status)
        color.Red(msg)
    }else{
        msg :=fmt.Sprintf("Name:%s  | Port :%d | %s ",s.Name,s.Port,status)
        color.Green(msg)
    }

}




func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return color.Pink(message)  
}
