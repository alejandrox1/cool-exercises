package main

import (
	"fmt"
	"time"
)

func tick(val int, d time.Duration) <-chan int {
	c := make(chan int)

	go func() {
		time.Sleep(d)
		c <- val
	}()

	return c
}

func main() {
	for i := 0; i < 24; i++ {
		c := tick(i, 100*time.Millisecond)
		fmt.Println(<-c)
	}
}
