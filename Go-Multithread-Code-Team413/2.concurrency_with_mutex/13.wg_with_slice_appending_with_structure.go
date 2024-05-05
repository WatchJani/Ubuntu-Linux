package main

import (
	"fmt"
	"sync"
)

func main() {
	x := make([]string, 0, 4)	// length 0 | capacity 4
	x = append(x, "start")
	fmt.Printf("%T",x)	// []string

	s1 := struct {
		List 		[]string
		Mutex 		sync.Mutex
	}{}
	fmt.Println(s1)		// []string{[] {0 0}}
	fmt.Printf("%T\n",s1)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	//================================================================
	// Setup input variables in go routines
	//================================================================
	go func(s1 *struct { List []string; Mutex sync.Mutex }, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		s1.Mutex.Lock()
		s1.List = append(s1.List, "1", "2")
		s1.Mutex.Unlock()
	}(&s1,wg)

	go func(s1 *struct { List []string; Mutex sync.Mutex }, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		s1.Mutex.Lock()
		s1.List = append(s1.List, "3", "4")
		s1.Mutex.Unlock()
	}(&s1,wg)

	go func(s1 *struct { List []string; Mutex sync.Mutex }, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		s1.Mutex.Lock()
		s1.List = append(s1.List, "5", "6")
		s1.Mutex.Unlock()
	}(&s1,wg)

	go func(s1 *struct { List []string; Mutex sync.Mutex }, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()
		s1.Mutex.Lock()
		s1.List = append(s1.List, "7", "8")
		s1.Mutex.Unlock()
	}(&s1,wg)

	wg.Wait()
	fmt.Println(s1.List)	// [1 2 7 8 3 4 5 6]
}
// Result will always have elements from all 4 wait groups : [start 7 8 1 2 3 4 5 6]