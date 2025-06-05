package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	val, err := rdb.HGet(ctx, "helloworldworkflow-app||dapr.internal.dapr-tests.helloworldworkflow-app.workflow||51d33015-7f97-4208-b2ed-19338e591aef||history-000005", "data").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
