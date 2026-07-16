package main 
import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello Budddy!!")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page!")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact Us!")
}
func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Profile!")
}

func main(){
	http.HandleFunc("/",home)
	http.HandleFunc("/about",about)
	http.HandleFunc("/contact",contact)
	http.HandleFunc("/profile",profile)
	fmt.Println("Server started at http://localhost:8080")
	err:= http.ListenAndServe(":8080",nil)

	if err != nil{
		fmt.Println(err)
	}
}