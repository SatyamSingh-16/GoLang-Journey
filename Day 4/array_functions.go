package main
import "fmt"

func main(){
	nums := [3]int{10,20,30}
	modify(nums)
	fmt.Println(nums)
}

func modify(arr [3]int){
	arr[0] = 100
}