package main

import "fmt"

func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {

	fmt.Println(Equal(10, 20))

	fmt.Println(Equal("Go", "Go"))

	fmt.Println(Equal(true, false))

}