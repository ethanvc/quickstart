package main

import (
	"context"
	"fmt"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, "http://127.0.0.1:8999", &websocket.DialOptions{
		Subprotocols: []string{"echo"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 5000000; i++ {
		err = wsjson.Write(ctx, c, map[string]int{
			"i": i,
		})
		if err != nil {
			fmt.Println(err)
		}

		v := map[string]int{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("received", v)

		if v["i"] != i {
			fmt.Printf("expected %v but got %v\n", i, v)
		}
		time.Sleep(2 * time.Second)
	}

	err = c.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		fmt.Println(err)
	}
}
