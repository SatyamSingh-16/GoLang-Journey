package main 
import "fmt"

func main(){
	var value any

	value = 100

	number,ok := value.(int)

	fmt.Println(number)
	fmt.Println(ok)

	value = "ok"

	number,ok = value.(int)

	fmt.Println(number)
	fmt.Println(ok)
}