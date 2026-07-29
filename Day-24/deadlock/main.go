package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex

	mu.Lock()

	fmt.Println("Locked once")

	mu.Lock()

	fmt.Println("This will never execute")

}
