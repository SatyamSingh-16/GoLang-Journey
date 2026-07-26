package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello Buddy!")
}

// func main() {

//		go sayHello()
//		fmt.Println("Main Finished")
//	}
func main() {

	go sayHello()

	time.Sleep(5 * time.Second)

	fmt.Println("Main Finished")
}
