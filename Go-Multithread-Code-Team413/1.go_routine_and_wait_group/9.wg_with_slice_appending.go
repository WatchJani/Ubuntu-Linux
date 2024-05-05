package main

import (
	"fmt"
	"sync"
)

func main() {
	x := []string{"start"}

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
// Result 1 : [start 3 4]
// Result 2 : [start 7 8 3 4 5 6]
// Result 3 : [start 7 8 1 2 3 4 5 6]