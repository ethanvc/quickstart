package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// trigger connection pool timeout error.
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "3.3.3.3:10000",
		PoolSize: 1,
	})
	for i := 0; i < 3; i++ {
		go get(rdb, "0")
	}
	time.Sleep(10000 * time.Hour)
}

func get(rdb *redis.Client, key string) {
	c, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	v := rdb.Get(c, key)
	fmt.Println(v)
}
