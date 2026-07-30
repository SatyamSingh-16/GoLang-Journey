package main

import "fmt"

func generator(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			out <- n
		}
	}()
	return out
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
func main() {
	numbers := generator(1, 2, 3, 4)
	squares := square(numbers)
	printer(squares)
}
