package main
import "fmt"

func main(){
	a := []int{1,2,3,4}
	b := a
	b[2] = 100
	fmt.Println("a =",a)
	fmt.Println("b =",b)
}