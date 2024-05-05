package main

import "fmt"

// Go channels are isolated processes which help us to share data between go routines.
// In Go Channel we can store value of any data type, regardless if that type is primitive or complex.
// On creation of the channel, we will specify type which that channel can store.

// We have two types of channels: buffered and unbuffered.
// In buffered channels we can store specific number of objects which we specified when we created channel.
// In unbuffered we can store as many as we want objects of specific type.

func main() {

	// Here we created go channel to which we can send data from go routine.
	// This is creation of unbuffered channel, where channel can store boolean data type.
	c := make(chan bool)
	// We started function in async mode, with two parameters.
	// First parameter is <c> go channel and second is <string> "World".
	go waitAndSay(c, "World")

	// This function will be executed right away after we execute waitAndStay() function in async mode.
	fmt.Println("Hello")

	// We are sending value to <c> channel which is passed to WaitAndSay() function.
	c <- true

	fmt.Println(<-c) // We are here until we receive value! This is blocking function!
	// We print value from channel, which in this case will be <false> value from waitAndStay() function!
	// Output:
	// Hello
	// World
	// false
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b { 		// We are receiving value from c channel!
		fmt.Println(s) 		// When we receive we execute function, if b is true! In this case we will print "World".
	}
	c <- false 				// We are sending true value to the channel!
}

// channel <- 5
// myNumber <- channel
// fmt.Println(<-channel)

//
// <-time.After(time.Second*6)
/*
<Operation rules>:
OPERATION 				A nil channel 		A closed Channel 		ANot-Closed Non-nil channel
CLOSE 					panic				panic 					succeed to close
SEND VALUE TO			block for ever		panic 					block or succeed to send
RECEIVE VALUE FROM		block for ever		never block				block or succeed to receive
*/
