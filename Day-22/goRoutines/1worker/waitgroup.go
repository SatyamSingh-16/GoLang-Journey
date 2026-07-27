package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printNumbers(&wg)
	wg.Wait() // this line will block the current goroutine...until wg becomes 0
	fmt.Println("FINISHED")
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println((i))
	}
}
