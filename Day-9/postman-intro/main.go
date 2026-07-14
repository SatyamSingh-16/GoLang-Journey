package main 
import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello Buddyy!!!!!")
}

func main(){

	http.HandleFunc("/hello",HelloHandler)
	fmt.Println("Server Started")
	http.ListenAndServe(":8080",nil)
}