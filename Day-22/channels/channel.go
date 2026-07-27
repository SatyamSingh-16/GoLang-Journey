package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	value := <-ch
	//called Rendezvous means
	//two goroutines meet to exchange data
	fmt.Println(value)
}
