package model

import (
	// "fmt"
	"time"
	"net/http"
)

type Service struct {
	ID int
	UserID int
	Name string
	URL string
	Healthy bool
	StatusCode int
	Response_time int64 
	Checked_at time.Time
}

func NewService(name,url string)Service{
	service :=Service{Name:name,URL:url,Healthy:false}
	return service
}
func (s Service)CheckHealth()(bool,int){
	//implement the logic to check the health of the service
	resp,err:=http.Get(s.URL)
	if err != nil{
		return false,0
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200,resp.StatusCode
}