package main

import ("sync";"fmt")

var y  = 0

func GlobalIncrement(wg *sync.WaitGroup, ch chan bool) {
	ch <- true	// Emulation of mutex.Lock()
	y = y + 1
	<- ch		// Emulation of mutex.Unlock()
	wg.Done()
}

func main() {
	// Here we didn't specify mutex lock and instead of that we used Go channel as blocking mechanism.
	// We have 1000 go routines and each go routine is manipulated with one single buffered channel.
	ch,wg := make(chan bool, 1),sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go GlobalIncrement(&wg, ch)
	}
	// Here we waited 1000 go routines to finish their execution.
	wg.Wait()
	fmt.Println("Final value of y", y)
}