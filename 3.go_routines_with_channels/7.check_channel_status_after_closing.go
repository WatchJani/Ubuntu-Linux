package main

import (
	"fmt"
	"strconv"
)

func SayHelloMultipleTimes(c chan string, n int) {
	for i := 1; i <= n; i++ {
		c <- "Hello " + strconv.Itoa(i)
	}
	//After we send n times Hello string we close the channel!
	close(c)
}

func main() {
	d := make(chan string)
	go SayHelloMultipleTimes(d, 3) // After execution this channel is closed!

	for s := range d { // Iterate over channel to see values!
		fmt.Println(s)
	}
	// Output:
	// Hello 1
	// Hello 2
	// Hello 3

	// Check is our channel is closed!
	// Our channel is now closed! We can't receive message anymore!
	v, ok := <-d // _, ok := (<-stop)
	if v == "" {
		v = "nil"
	}
	fmt.Println("Channel close?", !ok) // Channel close? true
	fmt.Println("Value: ", v)          // Value:  nil
}
