package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
)

var shared uint64
var sharedCount uint64

func readwrite(n int,wg *sync.WaitGroup) {
	defer wg.Done()	// Defer execution will be executed at the end of implementation of current process.
	iterations := rand.Int()%1000 + 1000
	fmt.Println("Cycle number: "+strconv.Itoa(n))
	// Atomic Operations - process that cannot be interrupted - Like we have Lock() Mutex for Writing.
	// This is how we can completely control variables, without using Mutex.
	// Atomic functions lock variables declared as primitive data types, for execution specific function for specific purpose.
	atomic.AddUint64(&shared, uint64(iterations))
	atomic.AddUint64(&sharedCount, uint64(iterations))

	// You can try to put incrementing with <++> operator, without Atomic Functions, so you can see the difference.
	// shared++
	// sharedCount++
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go readwrite(i,&wg)
	}
	wg.Wait()
	fmt.Println("Value of shared int =", shared)
	fmt.Println("Value of shared count =", sharedCount)
}

// Result without Atomic Functions:

// 				value of shared int = 70916
// 				value of shared count = 150611

// Many of shared++ operations were interrupted!
// When we change <shared++> to <atomic.AddUint64(&shared, 1)>, everything will be in order:

// Result:
// 				value of shared int = 150611
// 				value of shared count = 150611
