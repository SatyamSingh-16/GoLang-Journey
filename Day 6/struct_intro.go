package main 
import "fmt"

type Student struct{
	Name string
	Age int 
	CGPA float64
}

func main(){
	s1:= Student{
		Name: "Satyam",
		Age: 21,
		CGPA: 8.7,
	}
	fmt.Println(s1)
	fmt.Println(s1.Name)
	fmt.Println(s1.Age)
	fmt.Println(s1.CGPA)
}