package main

import "fmt"

func main() {
	c := make(chan int) // an unbuffered channel

	go func() {
		x := <- c // blocking here until a value is received.		> Step 2
		c <- x*x  // blocking here until the result is sent.		> Step 3
	}()

	c <- 3   // blocking here until the value is sent.				> Step 1
	y := <-c // blocking here until the result is received.			> Step 4
	fmt.Println(y) // 9
}
