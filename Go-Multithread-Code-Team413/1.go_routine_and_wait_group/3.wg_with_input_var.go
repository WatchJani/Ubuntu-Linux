package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = sync.WaitGroup{}
	var msg = "Hello"
	wg.Add(1)

	// Here we started new anonymous function as go routine, and inside of routine we change value of global variable.
	// Only in this case, we have input variable, which we want to pass to go routine.
	go func(msg string) {
		fmt.Println(msg) // Hello
		wg.Done()

	}(msg)
	// Here after function implementation, inside of () brackets, we can pass specific variable or value which we want.
	// In this case, we implemented <msg> variable as input variable of go routine.
	// Value which will be printed is "Hello" and not "Goodbye". Why?
	// Because at point when we passed <msg> variable to go routine, variable <msg> didn't change yet value.
	// At that point value of <msg> variable was still "Hello".

	msg = "Goodbye"
	wg.Wait()
}
