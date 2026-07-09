package main 
import (
	"errors"
	"fmt"
)

func main(){
	err := login(false)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("login successful")
}

func login(success bool) error {
	if !success {
		return errors.New("invalid credentials")
	}
	return nil
}