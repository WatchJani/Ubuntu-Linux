package main

import ("sync";"fmt")

// With mutex we can now lock or unlock something!
// This process allow us to read as many time we can get! 
		// But we can write one at the time!

// Infinite number of readers and one writer!

//Lock(): only one go routine read/write at a time by acquiring the lock.
//RLock(): multiple go routine can read(not write) at a time by acquiring the lock.
//Locker(): lock and unlock by "key"

//Remember we have classic Mutex, and RWMutex > 
		//Difference is that in Mutex we don't have RLock!

var x = 0
func RisePower(wg *sync.WaitGroup, m *sync.RWMutex) {
	m.Lock()	// We lock all other process so they can't be executed!
	x = x + 1	// Only one process can modify x at the present!
	m.Unlock()	// After unlock, first go routine that take access of x variable, can start execution!
	wg.Done()
}
// If we don't put Lock with Mutex, there will also mess, even we have wait group!
// When we have Lock(), only one go routine can access global variable!

func main() {
	var w sync.WaitGroup
	var m sync.RWMutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go RisePower(&w, &m)		// Here we will start 1000 go routines.
	}
	w.Wait()
	fmt.Println("Final value of x is", x)	// Final value of x is 1000
	fmt.Println("Checkpoint 1 Done!")

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go RisePower(&w, &m)		// Here we will start 1000 go routines.
	}

	w.Wait()
	fmt.Println("Final value of x is", x)	// Final value of x is 2000
	fmt.Println("Checkpoint 2 Done!")
}