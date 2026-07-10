package main

import "fmt"

func update(nums []int) {

	nums[1] = 500
	fmt.Println(nums)

}

func main() {

	arr := []int{10,20,30}

	update(arr)

	fmt.Println(arr)

}