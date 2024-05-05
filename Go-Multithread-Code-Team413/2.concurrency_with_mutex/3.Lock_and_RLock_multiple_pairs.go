package main

import ("sync";"fmt";"time")

// Lock(): only one go routine read/write at a time by acquiring the lock.
// RLock(): multiple go routine can read(not write) at a time by acquiring the lock.

func main() {
	lock,b := sync.RWMutex{},map[string]int{"0":0}

	// Here we have 4 different anonymous functions executed like go routines.
	// In first pair, we have RLock() and then Lock() as full lock.
	// In second pair, we have Lock() and then RLock().

	go func(i int) {		// If some is also reading, he can access data!
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["0"])
		time.Sleep(time.Second*2)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(1)

	go func(i int) {		// No one else can Read or Write data!
		lock.Lock()
		b["2"] = i
		fmt.Printf("Lock: from go routine %d: b = %d\n",i, b["2"])
		time.Sleep(time.Second*3)
		fmt.Printf("Lock: from go routine %d: lock released\n",i)
		lock.Unlock()
	}(2)

	//Our first two go routines last 5 second+couple ms.
	// That's why we will setup sleep time slightly more then that.
	time.Sleep(time.Second*5+time.Millisecond*100)
	// Here first pair of go routines finished execution.
	fmt.Println("=====================================================")

	go func(i int) {
		lock.Lock()
		b["3"] = i
		fmt.Printf("Lock: from go routine %d: b = %d\n",i, b["3"])
		time.Sleep(time.Second*3)
		fmt.Printf("Lock: from go routine %d: lock released\n",i)
		lock.Unlock()
	}(3)

	go func(i int) {
		lock.RLock()
		fmt.Printf("RLock: from go routine %d: b = %d\n",i, b["3"])
		time.Sleep(time.Second*3)
		fmt.Printf("RLock: from go routine %d: lock released\n",i)
		lock.RUnlock()
	}(4)

	// We can put here 5 seconds sleep!
	// In that case we won't see RLock() release method from last go routine. Why?
	// At that moment no one will wait for that go routine, and that routine won't give us response.
	time.Sleep(time.Second*7)
}
/* Here output is depending on which go routine will step up first into line for execution.
That's why we have multiple choices of output, and waiting.
We can also change sleeping time inside of go routines and at the end of application, so we can see the difference.

======================Possible Output 1==========================================
Lock: from go routine 2: b = 2
Lock: from go routine 2: lock released
RLock: from go routine 1: b = 0
RLock: from go routine 1: lock released
================================================================
RLock: from go routine 4: b = 0
RLock: from go routine 4: lock released
Lock: from go routine 3: b = 3
Lock: from go routine 3: lock released

======================Possible Output 2==========================================
RLock: from go routine 1: b = 0
RLock: from go routine 1: lock released
Lock: from go routine 2: b = 2
Lock: from go routine 2: lock released
================================================================
Lock: from go routine 3: b = 3
Lock: from go routine 3: lock released
RLock: from go routine 4: b = 3
RLock: from go routine 4: lock released
*/