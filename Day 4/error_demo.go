package main 
import (
	"fmt"
	"errors"
)

func main(){
	err := errors.New("Database connection failed!")
	fmt.Println(err)
	fmt.Println(err.Error())
}