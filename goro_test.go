package tcp_ip

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	c := make(chan int)

	go channel2(c)
	go channel1(c)

	fmt.Scanln()
}

func channel1(ch chan<- int) {
	ch <- 10
	fmt.Println("done1")
}

func channel2(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("done2")
}

// 송신용
func sendToChannel(ch chan<- string) {
	ch <- "hello"
}

// 수신용
func receiveFromChannel(ch <-chan string) {
	result := <-ch
	fmt.Println(result)
}
