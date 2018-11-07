package main

import (
	"fmt"
	"time"
)

func player(table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Println(ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	var Ball int

	table := make(chan int)

	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	fmt.Printf("Final value: %v\n", <-table)
}
