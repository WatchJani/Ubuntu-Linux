package main

import (
	"fmt"
	"sync"
)

var salutation = "Hello world!"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Here we started new anonymous function as go routine, and inside of routine we change value of global variable.
	go func() {
		salutation = "Welcome to Go world!"
		wg.Done()
	}()

	wg.Wait()
	// At this point, we finished with our WaitGroup, and global variable at this point has new value.
	fmt.Println(salutation) // Printed: Welcome to Go world!
}
