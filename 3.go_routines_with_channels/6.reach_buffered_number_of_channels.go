package main

import "fmt"

func main() {
	c := make(chan string, 2)

	trySend := func(v string) string {
		// When we manipulate over go channels we can create select loop.
		// This loop can separate for us different stages, when we receiving and sending values.
		// We can also check if at that point, that channel is empty or full if channel is buffered.
		select {
			case c <- v:
				return "Value <"+v+"> has been sent to <c> channel."
			default:
				// We will enter in default state, if <c> channel is full.
				return "We received new value: <"+v+">. Channel <c> is full. We can't write anything to <c> channel at the point."
		}
	}

	tryReceive := func() string {
		select {
			case v := <-c: return v
			default:
				// We will enter in default state, if <c> channel is empty.
				return "Channel <c> is empty. We don't have anything to read at the point."
		}
	}

	fmt.Println(trySend("Hello!"))
	fmt.Println(trySend("Hi!"))
	fmt.Println(trySend("Bye!")) 		// fail to send, but will not blocked.
	// Here with tryReceive() function we will take values from <c> channel.
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	fmt.Println(tryReceive()) // Channel <c> is empty. We don't have anything to read at the point.

	/* Output:
	Value <Hello!> has been sent to <c> channel.
	Value <Hi!> has been sent to <c> channel.
	Channel <c> is full. We can't write anything to <c> channel at the point.
	Hello!
	Hi!
	Channel <c> is empty. We don't have anything to read at the point.
	 */
}
