package tcp_ip

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	// 1
	go func() {
		defer func() { done <- struct{}{} }()
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			// 3
			fmt.Println("second goroutine")
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				fmt.Println("buf start")
				buf := make([]byte, 1024)
				for {
					//4
					n, err := c.Read(buf)
					fmt.Println(n)
					fmt.Println(err)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:56860")
	fmt.Println(conn)
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
	<-done
	listener.Close()
	<-done

}
