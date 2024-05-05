package main

import "fmt"

var chans = []chan int {make(chan int),make(chan int),make(chan int),make(chan int)}
var chans2 [5]chan int

func Run() {
	chans[0] <- 1	// Value 1 is sent to channel with index place 0 in <chans> list.
	chans[1] <- 2	// Value 1 is sent to channel with index place 1 in <chans> list.

	chans2[0] <- 3	// Value 3 is sent to channel with index place 0 in <chans2> list.
	chans2[1] <- 4	// Value 4 is sent to channel with index place 1 in <chans2> list.
}

func main(){
	fmt.Println(chans2)				// [<nil> <nil> <nil> <nil> <nil>]
	for i := range chans2 {			// We must fill channels with specific form.
		chans2[i] = make(chan int)	// In <chans2> list we will assign 5 new channels with with form to receive int data type.
	}
	fmt.Println(chans2)				// [0xc00001e240 0xc00001e2a0 0xc00001e300 0xc00001e360 0xc00001e3c0]
	go Run()
	obj,_ := <- chans[0]
	obj2,_ := <- chans[1]
	fmt.Println(obj,obj2)			// 1 2

	obj3,_ := <- chans2[0]
	obj4,_ := <- chans2[1]
	fmt.Println(obj3,obj4)			// 3 4
}

