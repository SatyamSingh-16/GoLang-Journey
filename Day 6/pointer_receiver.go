package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) Birthday() {
	s.Age++
}

func main() {

	student := Student{
		Name: "Satyam",
		Age: 21,
	}

	student.Birthday()

	fmt.Println(student.Age)

}