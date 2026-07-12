package main
import "fmt"

func Identity[T any](value T) T{
	return value
}

func main(){
	a := Identity(50)
	v:=Identity("Satyam")
	c:=Identity(true)

	fmt.Println(a)
	fmt.Println(v)
	fmt.Println(c)
}