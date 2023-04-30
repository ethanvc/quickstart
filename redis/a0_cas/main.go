package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
)

type CacheValue struct {
	Version int
	Data    int
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func main() {
	c := context.Background()
	val, err := getFromCache(c, "0")
	if err == redis.Nil {
		val = &CacheValue{}
	} else if err != nil {
		return
	}
	val.Data++
	err = writeCache(c, "0", val.Version, val)
	if err != nil {
		return
	}
}

func getFromCache(c context.Context, key string) (*CacheValue, error) {
	data, err := rdb.Get(c, key).Bytes()
	if err != nil {
		return nil, err
	}
	var obj CacheValue
	json.Unmarshal(data, &obj)
	return &obj, nil
}

var ErrVersionChanged = errors.New("DataVersionChanged")

func writeCache(c context.Context, key string, version int, val *CacheValue) error {
	err := rdb.Watch(c, func(tx *redis.Tx) error {
		oldVal, err := getFromCache(c, key)
		if err == redis.Nil {
			oldVal = &CacheValue{}
		} else if err != nil {
			return err
		}
		if oldVal.Version != version {
			return ErrVersionChanged
		}
		val.Version = version + 1
		_, err = tx.TxPipelined(c, func(pipe redis.Pipeliner) error {
			buf, _ := json.Marshal(val)
			pipe.Set(c, key, buf, 0)
			return nil
		})
		return err
	}, key)
	return err
}
