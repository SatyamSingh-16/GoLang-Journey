package main

import "fmt"

func main() {

	arr := [5]int{10,20,30,40,50}
	modify(arr)
	fmt.Println(arr)
	// s := arr[:3]

	// fmt.Println("Before")
	// fmt.Println(arr)
	// fmt.Println(s)

	// s = append(s,100)

	// fmt.Println("After")
	// fmt.Println(arr)
	// fmt.Println(s)

}
func modify(arr [5]int) {
	arr[1] = 40
	fmt.Println(arr)
}