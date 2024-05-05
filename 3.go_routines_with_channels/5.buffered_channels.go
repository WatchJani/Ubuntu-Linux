package main

import "fmt"

func main(){
	// Buffered Channels-Declare:
	//								ch:=make(chan <type>, <number of buffered channels>)
	//								ch:=make(chan int, 100)
	// Here we defined buffered channel where we can store only two objects of <string> type.
	// Senders only has been blocked when the buffer is full! 	>>> All channels are busy!
	// Receivers block when the buffer is empty!			   	>>> We closed all the channels!
	ch:=make(chan string,2)

	ch<-"Hello"
	ch<-"World"
	// If we try to send new value after we send these two values, we will get error:
	// 				ERROR: We have deadlock, and we dont have space for 3rd value!

	fmt.Println(<-ch,<-ch)	// Hello World
}
