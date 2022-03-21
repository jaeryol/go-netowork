package tcp_ip

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"syscall"
	"testing"
	"time"
)

type HttpHandler struct{} // HttpHandler의 method를 구현한다.

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}

func TestDialContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	//go func() {
	//	http.ListenAndServe("127.0.0.1:3001", HttpHandler{})
	//}()

	go func() {
		defer func() {
			sync <- struct{}{}
		}()
		var d net.Dialer
		d.Control = func(ctx, address string, c syscall.RawConn) error {
			time.Sleep(5 * time.Second)
			fmt.Println("call? , ", c)
			return nil
		}
		fmt.Println("start")
		conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
		fmt.Println("conn ", conn, " / ", err)
		if err != nil {
			t.Log("errrr > ", err)
			return
		}
		conn.Close()
		t.Error("connection did not time out")
	}()
	fmt.Println("before cancel()")
	cancel()
	fmt.Println("after cancel()")
	<-sync
	if ctx.Err() != context.Canceled {
		t.Errorf("expected canceled context; actual: %v", ctx.Err())
	}

}
