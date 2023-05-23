package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	// key exist or not, no err return.
	cmd := rdb.TTL(context.Background(), "33")
	err := cmd.Err()
	fmt.Println(cmd, err)
}
