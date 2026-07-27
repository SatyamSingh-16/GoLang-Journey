package main

import (
	"fmt"
	"time"
)

func main() {
	paymentCh := make(chan string)
	select {
	case payment := <-paymentCh:
		fmt.Println(payment)
	case <-time.After(3 * time.Second):
		fmt.Println("Payment Timed Out")
	}

}
