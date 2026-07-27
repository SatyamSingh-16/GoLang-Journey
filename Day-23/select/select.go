package main

import (
	"fmt"
	"time"
)

func main() {

	studentCh := make(chan string)
	teacherCh := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		studentCh <- "Student Data Ready"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		teacherCh <- "Teacher Data Ready"
	}()
	select {

	case message := <-studentCh:

		fmt.Println(message)

	case message := <-teacherCh:

		fmt.Println(message)

	}
}
