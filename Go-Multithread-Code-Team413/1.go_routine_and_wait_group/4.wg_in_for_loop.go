package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		// If we have multiple processes which we need to execute inside of <for> loop,
		// we can start multiple go routines under <for> loop.

		// On each iteration, we can access WaitGroup, call Add() function, and increase number of waiting processes by 1.
		// In this case, each iteration will increase by 1 WaitGroup, so at the end we will wait 6 processes to end their implementation.
		wg.Add(1)
		go func(i int) {
			// Inside of
			defer wg.Done()
			// We have 3 checkpoints: after 1,2 and 3 seconds.
			// Maybe in first we have 2, in second we have 1, and in the last we have 3 processes!
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for ", i)
		}(i) 	// Here variable <i> will be input variable for our go routine.
	}

	wg.Wait()
	fmt.Println("All jobs has been finished!")
	/*
		Work done for  0
		Work done for  1
		Work done for  3
		Work done for  2
		Work done for  5
		Work done for  4
		All jobs has been finished!
	*/
}
