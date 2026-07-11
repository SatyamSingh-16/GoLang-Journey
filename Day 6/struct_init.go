package main
import "fmt"

type Student struct {
	Name string
	Age int 
	CGPA float64
}

func NewStudent(name string,age int) Student{
	return Student{
		Name: name,
		Age: age,
		CGPA: 8.0,
	}
}

func main(){

	s1:= NewStudent("Satyam",21)
	fmt.Println(s1)
	
	s2:= Student{
		Name: "Rahul",
		Age: 22,
		CGPA: 8.9,
	}

	fmt.Println(s2)
}