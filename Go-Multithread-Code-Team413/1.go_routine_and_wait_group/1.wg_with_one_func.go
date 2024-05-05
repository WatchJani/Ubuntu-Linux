package main

import (
	"fmt"
	"sync"
)

func main() {
	// Wait group is structure where we define, how many processes will create one checkpoint.
	// We declare new variable as WaitGroup and then with Add() function add number of the processes.
	var wg sync.WaitGroup
	wg.Add(1)
	// In this case WaitGroup will wait only one process in group.
	// As we know we can define variable as function, and then execute that function.
	sayHello := func() {
		fmt.Println("Hello people!")
		// Here we access WaitGroup and tell the group that this process has finished with implementation.
		wg.Done()
	}
	// In wait group we need to specify all processes which are defined and started as Go Routines.
	// If some function is started as Go Routine, we started that process async.
	// That means that our application won't wait that process to finish whole implementation.
	// Instead of that, application will start the process and immediately continue with next process.
	// To define specific process as Go Routine, we need to specify reserved word <go> in from of function execution.

	go sayHello() // Release go routines as async function!

	// On the place, where we want to create checkpoint, and inform our WaitGroup that all processes are finished,
	// we will access WaitGroup and execute Wait().
	wg.Wait()
}
