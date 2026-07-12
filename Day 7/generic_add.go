package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

type Age int

func main() {

	fmt.Println(Add(10,20.9))

	fmt.Println(Add(3.5,2.5))

	var x Age = 25

	fmt.Println(Add(x,x))

}