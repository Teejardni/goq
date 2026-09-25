package main

import (
	"github.com/redis/go-redis/v9"

	"github.com/Teejardni/goq/queue"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_ = queue.NewRedisQueue(client)
}
