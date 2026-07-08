package main 
import "fmt"

func main(){
	sum,_ := calculate(10,15)
	fmt.Println(sum)
	// fmt.Println(product)
}

func calculate(a,b int) (int,int){
	sum := a + b 
	product := a*b 

	return sum,product
}