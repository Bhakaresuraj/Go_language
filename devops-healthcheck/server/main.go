package main
import(
	"net/http"
)

func checkHandler(w http.ResponseWriter	,r * http.Request){
	
}

func main(){
	http.HandleFunc("/check",checkHandler)
	http.ListenAndServe(":3000",nil);
}