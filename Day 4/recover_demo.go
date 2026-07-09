package main

import "fmt"

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}

	}()

	fmt.Println("Program Started")

	panic("Unexpected Error")

	fmt.Println("Program Finished")

}