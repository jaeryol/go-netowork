package tcp_ip

import (
	"fmt"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	fmt.Printf("listener %s ", listener.Addr())
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		fmt.Println("final call")
		_ = listener.Close()
	}()
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	go func(c net.Conn) {
	//		defer c.Close()
	//	}(conn)
	//}
	t.Logf("bound to %q", listener.Addr())
}
