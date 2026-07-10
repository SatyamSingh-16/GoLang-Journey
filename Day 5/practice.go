package main 
import "fmt"

func main(){
	b := make([]int, 5)

	c := make([]int, 5, 10)

	fmt.Println(b, len(b), cap(b))

	fmt.Println(c, len(c), cap(c))
}