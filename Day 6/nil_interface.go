package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct{}

func (d *Dog) Sound() {
	fmt.Println("Woof!")
}

func main() {

	var a Animal

	fmt.Println(a)

	// a.Sound()

}