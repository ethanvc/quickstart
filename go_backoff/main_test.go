package go_backoff

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	bo := backoff.NewTicker(backoff.NewExponentialBackOff())
	curT := time.Now()
	for {
		select {
		case <-bo.C:
		}
		fmt.Println(time.Now().Sub(curT).Milliseconds())
		curT = time.Now()
	}
}
