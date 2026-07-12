package main 
import "fmt"

func Descrive(value any){

	switch v:= value.(type){
	case int:
		fmt.Println("Integer:",v)
	case string:
		fmt.Println("String:",v)
	case bool:
		fmt.Println("Boolean:",v)
	default:
		fmt.Println("Unknown")
	}
}

func main(){
	Descrive(10)
	Descrive(true)
	Descrive("Satyam")
	Descrive(3.13)
}