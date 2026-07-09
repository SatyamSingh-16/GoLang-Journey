package main 
import "fmt"

func main(){
	defer fmt.Println("Cleaning Resources")
	fmt.Println("Application Started")
	panic("Unexpected failure")
	fmt.Println("Application Finished")
}