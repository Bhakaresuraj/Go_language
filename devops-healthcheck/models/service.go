package service

type Service struct{
    Name string
    Port int
    Healthy bool 
    desc string
}

func NewService(name string ,port int) Service{
    isHealthy := ValidatePort(port)
    service :=Service{Name:name, Port:port,Healthhy:isHealthy}
    return service
} 

func ValidatePort(port int)bool{
    if port <=0 ||port >=65535{
        return false
    }else{
        return true
    }

}
