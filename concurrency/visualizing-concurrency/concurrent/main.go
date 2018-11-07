package main

import "fmt"

// Generate sends the sequence 2, 3, 4, ..., to the channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter copies the moves from the 'in' channel to the 'out' channel, removing
// those diviseible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)

	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)

		out := make(chan int)
		go Filter(ch, out, prime)

		ch = out
	}
}
