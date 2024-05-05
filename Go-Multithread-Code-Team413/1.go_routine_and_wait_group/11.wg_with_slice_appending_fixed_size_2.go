package main

import (
	"fmt"
	"sync"
)

func main() {
	x := make([]string, 0, 4)	// length 0 | capacity 4
	x = append(x, "start")
	fmt.Printf("%T",x)	// []string

	wg := &sync.WaitGroup{}
	wg.Add(4)
	//================================================================
	// Setup input variables in go routines
	//================================================================
	go func(list *[]string, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		*list = append(*list, "1", "2")
	}(&x,wg)
	go func(list *[]string, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		*list = append(*list, "3", "4")
	}(&x,wg)
	go func(list *[]string, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		*list = append(*list, "5", "6")
	}(&x,wg)
	go func(list *[]string, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		*list = append(*list, "7", "8")
	}(&x,wg)
	wg.Wait()
	fmt.Println(x)
}
// Result will always have elements from all 4 wait groups : [start 7 8 1 2 3 4 5 6]