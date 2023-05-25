package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"nhooyr.io/websocket"
)

func main() {
	http.ListenAndServe("127.0.0.1:8999", http.HandlerFunc(serveHttp))
}

func serveHttp(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{Subprotocols: []string{"echo"}})
	if err != nil {
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
	if c.Subprotocol() != "echo" {
		c.Close(websocket.StatusPolicyViolation, "client must speak the echo sub protocol")
		return
	}
	for {
		err = echo(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func echo(c context.Context, conn *websocket.Conn) error {

	typ, r, err := conn.Reader(c)
	if err != nil {
		return err
	}

	w, err := conn.Writer(c, typ)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}
	return nil
}
