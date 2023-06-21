package main

import (
	"fmt"
	"time"
)

type User struct {
	Uid  int64
	Flag int64
}

func main() {
	m := make(map[int64]*User)
	fmt.Printf("start at %s\n", time.Now().Format(time.RFC3339Nano))
	s := time.Now()
	for i := 0; i < 50000000; i++ {
		m[int64(i)] = &User{
			Uid:  int64(i),
			Flag: int64(i),
		}
		if time.Now().Sub(s) > 10*time.Second {
			fmt.Printf("index %d\n", i)
			s = time.Now()
		}
	}
	fmt.Printf("end at %s\n", time.Now().Format(time.RFC3339Nano))
	for {
		fmt.Printf("loop at %s\n", time.Now().Format(time.RFC3339Nano))
		time.Sleep(time.Minute)
	}
}
