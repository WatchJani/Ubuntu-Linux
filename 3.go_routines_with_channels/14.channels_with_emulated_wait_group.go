package main

import (
	"fmt"
	"sync"
	"time"
)

type CounterMC struct{
	N 	int
	sync.RWMutex
}
func (c *CounterMC) IncreaseWaitingValues(){
	c.Lock()
	c.N++
	c.Unlock()
}
func (c *CounterMC) GetWaitingValues() int {
	c.RLock()
	g := c.N
	c.RUnlock()
	return g
}
//============================================================================
func main() {
	messages := make(chan int)
	// We passed the same channel to 3 different go routines.
	go func(ch chan int) {
		time.Sleep(time.Second * 3)
		ch <- 1
	}(messages)
	go func(ch chan int) {
		time.Sleep(time.Second * 2)
		ch <- 2
	}(messages)
	go func(ch chan int) {
		time.Sleep(time.Second * 1)
		ch <- 3
	}(messages)
	cmc := &CounterMC{}
	for m := range messages {
		fmt.Println(m)
		cmc.IncreaseWaitingValues()
		if cmc.GetWaitingValues() == 3 {
			break
		}
	}
}
