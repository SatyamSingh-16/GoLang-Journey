package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	select {

	case <-ctx.Done():

		fmt.Println("Worker cancelled")

		return

	}

}

func main() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	go worker(ctx)

	time.Sleep(5 * time.Second)

}
