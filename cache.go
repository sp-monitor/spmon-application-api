package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func writeCache(moduleData ClientData) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
		DB: 0,
	})

	ctx := context.Background()

	err := client.Publish(ctx, "module_data", fmt.Sprintf("%v", moduleData)).Err()
	if err != nil {
		fmt.Println(err)
	}
}
