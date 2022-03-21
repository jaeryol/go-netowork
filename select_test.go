package tcp_ip

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func longFunc() string {
	<-time.After(time.Second * 3) // long running job
	return "Success"
}

func TestHelloWorld(t *testing.T) {
	now := time.Now()
	fmt.Println("now > ", now)
	ctx, _ := context.WithDeadline(
		context.Background(),
		now.Add(2*time.Second),
	)

	done := make(chan string)
	go func() {
		done <- longFunc()
	}()
	fmt.Println("start")
	select {
	case result := <-done:
		fmt.Println(result, ", ", nil)
	case <-ctx.Done():
		fmt.Println("Fail, ", ctx.Err())
	}
	fmt.Println("wait")
}
