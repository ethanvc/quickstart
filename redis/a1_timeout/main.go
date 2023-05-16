package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// trigger connection pool timeout error.
// docker run --rm -it -p 10000:6379 redis
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "3.3.3.3:10000",
		PoolSize: 1,
	})
	for i := 0; i < 1; i++ {
		go get(rdb, "0")
	}
	time.Sleep(10000 * time.Hour)
}

func get(rdb *redis.Client, key string) {
	c, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()
	v := rdb.Get(c, key)
	fmt.Println(v)
}
