package main

import (
	"fmt"
	"sync"
)

func gen5(nums ...int) chan int{

	out:=make(chan int)
	go func() {
		for _,n:=range nums{
			out<-n
		}
		close(out)
	}()
	return out
}

func sq5(in chan int) chan int{
	out:=make(chan int)
	go func() {
		for n:=range in{
			out<-n*n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int{
	var wg sync.WaitGroup
	out:=make(chan int)
	// Start an output go routines for each input channel in cs.
	// Output copies values from c to out until c is closed, then calls wg.Done.
	// Here define method which will be our go routine in definition below.
	output:= func(c <-chan int) {
		for n:=range c{
			out<-n
		}
		wg.Done()
	}
	//Add delta to WaitGroup > how many channels do we have.
	wg.Add(len(cs))
	for _,c:=range cs {
		go output(c)
	}
	//Start a goroutine to close out once all the outputs are done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main(){
	in:=gen5(2,3)
	c1:=sq5(in)
	c2:=sq5(in)
	for n:= range merge(c1,c2){
		fmt.Println(n)
		// Output: 4 then 9, or 9 then 4
	}
}

/*
FAN OUT
Multiple 6.functions reading from the same channel until that channel is closed

FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input 160.channels_medium onto a single channel that's closed when all
the inputs are closed.

PATTERN
There's a pattern to our pipeline 6.functions:
	>>stages close their outbound 160.channels_medium when all send operations are done
	>>stages keep receiving values from inbound 160.channels_medium until those 160.channels_medium are closed

source:
https://blog:golang.org/pipelines
*/
