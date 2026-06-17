package main

import (
	"fmt"
	"net/http"
	"os"
	
)

func checkHandler(w http.ResponseWriter	,r * http.Request){
	
}

func main(){
	http.HandleFunc("/check",checkHandler)
	fmt.Print(os.Getenv("DATABASE_URL"))
	http.ListenAndServe(":5000",nil);
}