package main

import (
	"fmt"
	"net"
	"time"
)

func handler(c net.Conn, ch chan string) {
	ch <- c.RemoteAddr().String()

	time.Sleep(100 * time.Millisecond)

	c.Write([]byte("hi!"))
	c.Close()
}

func logger(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		fmt.Println(<-results)
	}
}

// pool starts n loggers. All the messages from ch will be sent to wch,
// registering the leght of these messages.
func pool(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)

	for i := 0; i < n; i++ {
		go logger(wch, results)
	}

	go parse(results)
	for {
		addr := <-ch
		wch <- len(addr)
	}
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go handler(c, ch)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	go pool(ch, 4)
	go server(l, ch)

	time.Sleep(15 * time.Second)
}
