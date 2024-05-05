package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg1, wg2 sync.WaitGroup
	for _, salutation := range []string{"Hello", "Welcome", "Good Day"} {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			fmt.Printf("<%v> ", salutation) // <Good Day> <Good Day> <Good Day>
			// Here we printed 3 times the same value, because all 3 go routines, started from the same point.
		}()
	}
	wg1.Wait()	// Here we will wait all routines from WaitGroup1, so we can start Checkpoint 2.
	fmt.Println("")

	for _, salutation := range []string{"Hello", "Welcome", "Good Day"} {
		wg2.Add(1)
		go func(salutation string) { // Put input variable in go func()
			defer wg2.Done()
			fmt.Printf("<%v> ", salutation) // <Good Day> <Hello> <Welcome>
		}(salutation)
		// Here we printed 3 times different values like we have in iteration.
		// All 3 go routines, started from different starting point, because we passed value from slice to go routine as input variable.
	}
	wg2.Wait()	// Here we will wait all routines from WaitGroup2.
}
