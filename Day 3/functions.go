package main
import "fmt"

func main(){
	resultAdd := add(15,25)
	resultMultiply := multiply(15,25)
	fmt.Println(resultAdd)
	fmt.Println(resultMultiply)
}

func add(a,b int) int{
	return a + b
}

func multiply(a,b int) int{
	return a * b
}