package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Here we have two different function with totally different implementation, but they belong to the same WaitGroup.
	// We will wait both of them to finish their implementation, so we can inform WaitGroup that all processes are done their jobs.
	// Also, in this case, we passes WaitGroup as input variable.
	// Inside of functions implementation, we access that WaitGroup and call Done() function.
	// That's how we informed WaitGroup that process has finished his implementation.

	go CollectPower(41, &wg)
	go DividePower(19432134, 41, &wg)

	wg.Wait()

	fmt.Println("Both functions has finished their work!")
	// 473954
	// 1684
	// Both functions has finished their work!

}

func CollectPower(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := (a * a) ^ 5
	fmt.Printf("%v\n", result)
}

func DividePower(power int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	strength := power / b
	fmt.Printf("%v\n", strength)
}
