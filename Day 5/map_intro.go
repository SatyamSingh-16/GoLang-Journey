package main
import "fmt"

func main(){

	marks:= map[string]int{
		"Alics": 95,
		"Bob": 82,
		"Charlie": 99,
	}

	fmt.Println(marks)
	fmt.Println(marks["Bob"])
}