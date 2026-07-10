package main
import "fmt"

func main(){

	marks:= map[string]int{
		"Satyam": 100,
		"Shivansh": 99,
	}

	score,ok := marks["Satyam"]
	fmt.Println(score)
	fmt.Println(ok)

	if ok {
		fmt.Println("Score exist and is equal to",score)
	}else{
		fmt.Println("No user exist")
	}
}