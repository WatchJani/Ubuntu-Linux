package main

import ("sync";"fmt";"time")

//Lock(): only one go routine read/write at a time by acquiring the lock.
//RLock(): multiple go routine can read(not write) at a time by acquiring the lock.

func main() {

	a,b,lock := "WRITE","READ",sync.RWMutex{}

	for i := 1; i < 5; i++ {		// Here we have 4 Lock Mutex!
		go func(i int) {			// Here we write so Read can't be executed!
			lock.Lock()
			fmt.Printf("Lock: from go routine %d: a = %v\n",i, a)
			time.Sleep(time.Second)
			lock.Unlock()
		}(i)
	}

	for i := 11; i < 20; i++ {		// Here we have 9 Read Mutex executions!
		go func(i int) {
			lock.RLock()
			fmt.Printf("RLock: from go routine %d: b = %v\n",i, b)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second*5)
	// After these 5 seconds, we will have print, where at some point <WRITE> mutex will interrupt <READ> mutex,
	// it will lock all other processes and back all of them into line, so they need to wait for writing process to finish.
}
/*
Lock: from go routine 1: a = WRITE
RLock: from go routine 13: b = READ
RLock: from go routine 17: b = READ
RLock: from go routine 15: b = READ
RLock: from go routine 19: b = READ
RLock: from go routine 16: b = READ
RLock: from go routine 12: b = READ
RLock: from go routine 11: b = READ
RLock: from go routine 18: b = READ
RLock: from go routine 14: b = READ
Lock: from go routine 2: a = WRITE
Lock: from go routine 3: a = WRITE
Lock: from go routine 4: a = WRITE
*/