package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i	// Here we will send values from 0 to 9 to our <c> channel.

		}
		close(c)				// If we close the channel, we can read from it, but we can not add any more values.
	}()
	// If we closed the channel c, we can read all data from it by looping over range.
	// If we don't close the channel and we try to read channel values over range function, we will get:
	//					fatal error: all goroutines are asleep - deadlock!
	for n := range c {
		fmt.Println(n)
	}
}
