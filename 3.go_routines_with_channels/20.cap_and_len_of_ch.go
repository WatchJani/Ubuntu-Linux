package main

import "fmt"

func main() {
	c := make(chan int, 2) // Buffered channel
	c <- 3
	c <- 5
	close(c)	// Closed channel is not empty, we just can't trasfer data!

	//We have capacity and length of channel!
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok) // 3 true
	fmt.Println(len(c), cap(c)) // 1 2 > We get one value from channel!
							 	// 	   > Now that value is deleted from channel!
	x, ok = <-c
	fmt.Println(x, ok) // 5 true
	fmt.Println(len(c), cap(c)) // 0 2

	x, ok = <-c
	fmt.Println(x, ok) // 0 false		
}
