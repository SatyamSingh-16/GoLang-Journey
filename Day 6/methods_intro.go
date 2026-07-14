package main
import "fmt"


type Student struct{
	Name string
	Age int
}

func (s Student) Introduction(){
	fmt.Println("Hi, I'm",s.Name)
	fmt.Println("I'm",s.Age,"years old")
}
func (s *Student) Greeting() {
	s.Age++
}

func main(){
	s1:=Student{
		Name: "Satyam",
		Age: 30,
	}

	s1.Introduction()
	s1.Greeting()
	fmt.Println(s1.Age)
}