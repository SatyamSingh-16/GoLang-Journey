package main 
import "fmt"

func main(){
	arr:= [5]int{1,2,3,4,5}
	s:=arr[1:4]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}