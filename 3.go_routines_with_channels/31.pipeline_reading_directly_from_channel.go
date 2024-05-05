package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculate(msg int) <-chan int {
	c := make(chan int)
	value := msg
	go func() {
		for i:=0;;i++{
			c <- value
			value+=10
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
// We can write this also as: (chans ...<-chan string)
func getResults(input1, input2 <- chan int) <-chan int {
	c:= make(chan int)
	// Reading from two channels. Here we are sending values from <input1> and <input2> channels to <c> channel.
	go func() {
		for{
			c<- <-input1	// Read from one channel and send that value right away to another channel.
		}
	}()
	go func() {
		for{
			c<- <-input2	// Read from one channel and send that value right away to another channel.
		}
	}()
	// Sending directly to channel so we can read in above for loop immediately.
	return c
}

func main(){
	c := getResults(calculate(413), calculate(513))
	//We are waiting 10 values to come to this channel from 2 <calculate> functions with two different parameters.
	for i:=0;i<10;i++{
		fmt.Println(<-c)
	}
	fmt.Println("Calculations are finished, you can send data to Data Center...")
}
/*Output:
513
413
523
423
533
433
543
443
553
453
Calculations are finished, you can send data to Data Center...
 */
/*
             <FAN OUT>						<FAN IN>


                  ____________		____________
some channel	 |					____________|	some channel
_________________|____________		____________|____________________
				 |					____________|
				 |____________		____________|


>>multiple funcs reading from		>>multiple writing to
that channel until it's closed		the same channel

*/
