package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// docker run --rm -it -p 10000:6379 -name slow_redis redis
// sudo pfctl -f slow_redis.conf
// query if firewall well:
// sudo pfctl -sr
// enable fireall:
// sudo pfctl -e
// disable firewall:
// sudo pfctl -d
// query if enabled:
// sudo pfctl -s -info

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":10000",
		// PoolSize: 1,
	})
	for i := 0; i < 3; i++ {
		go get(rdb, "0")
	}
	time.Sleep(10000 * time.Hour)
}

func get(rdb *redis.Client, key string) {
	v := rdb.Get(context.Background(), key)
	fmt.Println(v)
}
