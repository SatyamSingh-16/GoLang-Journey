package main

import (
	"context"
	"fmt"
)

func main() {

	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	fmt.Println(parent)
	fmt.Println(ctx)

	cancel()
}
