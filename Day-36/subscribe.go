package day36

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func subscribeToOrders(
	ctx context.Context,
	client *redis.Client,
) error {

	pubsub := client.Subscribe(
		ctx,
		"order.created",
	)

	defer pubsub.Close()

	for {
		message, err := pubsub.ReceiveMessage(ctx)

		if err != nil {
			return err
		}

		fmt.Println(
			"Received:",
			message.Payload,
		)
	}
}
