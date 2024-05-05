package main

import (
	"fmt"
	"sync"
)

// Lock(): only one go routine read/write at a time by acquiring the lock.
// RLock(): multiple go routine can read(not write) at a time by acquiring the lock.

type SafeCounter struct {
	sync.Mutex
	// sync.Mutex || sync.RWMutex
	// Here mutex is embedded to struct.
	// This means that this Lock() or Unlock() are attached only to this structure.
	// If multiple go routine are reading this structure, or they want to write multiple values to attributes, we can use Mutex.
	// That's how we can manipulate over single object with multiple go routines.
	// We will have 100% of control and we will avoid the conflict between reading and writing.

	// We can also put Mutex like new attribute as new object inside of structure:
	// Mux 	   sync.Mutex
	ID int
}

// In each of these 3 functions, we can access Mutex inside of structure, and Lock() or Unlock() structure.
// In that case, we will execute all possible execution of functions connected to this structure.
// We won't have a conflict between reading and writing, but also between multiple writings.
func (sc *SafeCounter) Increment(y *int, g *sync.WaitGroup) {
	g.Add(1)
	sc.Lock()
	*y++
	sc.ID++
	fmt.Println("Inc")
	sc.Unlock()
	g.Done()
}

func (sc *SafeCounter) Decrement(y *int, g *sync.WaitGroup) {
	g.Add(1)
	sc.Lock()
	*y++
	sc.ID--
	fmt.Println("Dec")
	sc.Unlock()
	g.Done()
}

func (sc *SafeCounter) GetValue() int {
	sc.Lock()
	v := sc.ID
	sc.Unlock()
	return v
}

func main() {
	sc, p := new(SafeCounter), 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ { // We will run 200 proccesses and wait for them!
		go sc.Increment(&p, &wg)
		go sc.Decrement(&p, &wg)
	}
	wg.Wait()	// Here we will wait for 2000 go routines to finish their work.
	fmt.Printf("%v > %v > %v\n", sc.ID, sc.GetValue(), p) // 0 > 0 > 200
	// ID=0		GetValue()=0	  And we executed 200 proccesses!
}
