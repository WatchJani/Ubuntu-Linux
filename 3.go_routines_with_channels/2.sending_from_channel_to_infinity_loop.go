package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i	// Here we will send values from 0 to 9 to our <c> channel.

		}
		// close(c)				// If we close the channel, we can read from it, but we can not add any more values.
	}()

	go func() {
		for {
			fmt.Println(<-c)
			// Here in infinity for loop, we will take and read values from <c> channel, as long channel is active.
			// After we get value from channel, buffer of that channel is changed.
			// That means that we took first value which is in order from channel and after that value is removed from channel.
			// If we take next value in order from channel, after reading that value will be also removed from channel.
			// We will continue with reading until, we take all values from channel, and until channel is empty.
			// In this for loop we will read 10 different values which will be sent to <c> channel.
		}
	}()
	time.Sleep(time.Second)

	//
	for n := range c {				// If we closed the channel c, we can read all data from it by looping over range!
		fmt.Println(n)
	}
}
