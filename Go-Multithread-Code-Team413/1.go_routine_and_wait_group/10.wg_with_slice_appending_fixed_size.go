package main

import (
	"fmt"
	"sync"
)

func main() {
	x := make([]string, 0, 4)	// length 0 | capacity 4
	x = append(x, "start")
	fmt.Printf("%T",x)	// []string

	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		x = append(x, "1", "2")
	}()
	go func() {
		defer wg.Done()
		x = append(x, "3", "4")
	}()
	go func() {
		defer wg.Done()
		x = append(x, "5", "6")
	}()
	go func() {
		defer wg.Done()
		x = append(x, "7", "8")
	}()
	wg.Wait()
	fmt.Println(x)
}
// Result will always have elements from all 4 wait groups : [start 7 8 1 2 3 4 5 6]