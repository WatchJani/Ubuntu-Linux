package main

import (
	"fmt"
	"sync"
)

type User struct {
	Mtx  sync.RWMutex // Mutex as attribute inside of struct!
	Name string
}

func main() {
	p1 := User{}
	p1.Name = "Neo"
	fmt.Println(p1.Name, "- First Print")

	var wg sync.WaitGroup
	wg.Add(12)
	for i := 0; i < 10; i++ {
		// Here we will have 10 readings from 10 different go routines.
		// They will be intercepted by two go routines, when they execute their writings.
		// Those two functions will be both executed as SetName() functions in separated go routines.
		go p1.ReadName(&wg)
	}

	// Here we have two go routines, which will at some point, execute two writings to setup new value for name.
	// That means at some point, they will interrupt reading go routines, so they can modify Name with new value.
	go p1.SetName("Agent Smith", &wg)
	go p1.SetName("Agent 47", &wg)
	wg.Wait()
	fmt.Println(p1.Name, "- Last Print")
	fmt.Println("12 processes has been executed in WaitGroup! 2 writings and 10 readings!")
}

func (p *User) SetName(name string, g1 *sync.WaitGroup) {
	p.Mtx.Lock()
	p.Name = name
	fmt.Println("Write has been made! All reads are stoped! New Value="+name+"!")
	defer p.Mtx.Unlock()
	g1.Done()
}

func (p *User) ReadName(g1 *sync.WaitGroup) {
	p.Mtx.RLock()
	fmt.Println(p.Name)
	defer p.Mtx.RUnlock()
	g1.Done()
}

/*
Neo - First Print
Write has been made! All reads are stopped!					> At this point we have writing!
Agent 47													> After that reading go routines can continue with execution!
Agent 47
Agent 47
Agent 47
Agent 47
Write has been made! All reads are stopped! New Value=Agent Smith!		> At this point we have writing again!
Agent Smith																> After that the rest of reading go routines can continue with execution!
Agent Smith
Agent Smith
Agent Smith
Agent Smith
Agent Smith - Last Print
12 processes has been executed in WaitGroup! 2 writings and 10 readings!
*/

/*
Instruction	Goroutine 	1	Goroutine 				2	Bank Balance
1						Read balance ⇐ £50			£50
2						Read balance ⇐ £50			£50
3						Add £100 to balance			£50
4						Add £50 to balance			£50
5						Write balance ⇒ £150		£150
6						Write balance ⇒ £100		£100
*/
