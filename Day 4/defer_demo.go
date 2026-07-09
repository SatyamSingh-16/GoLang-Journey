package main 
import "fmt"

func main(){
	// fmt.Println("A")
	// defer fmt.Println("B")
	// fmt.Println("C")


	fmt.Println("Start")

	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

}