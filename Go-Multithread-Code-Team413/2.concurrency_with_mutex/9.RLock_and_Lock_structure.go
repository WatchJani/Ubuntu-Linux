package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Person struct {
	sync.RWMutex // Embedded mutex inside of struct!
	Name         string
}

func main() {
	p1 := &Person{}
	// Access mutex inside of structure, from anonymous function and Lock it.
	go func() {
		p1.Lock()
		p1.Name = "Neo"
		p1.Unlock()
	}()

	for i := 0; i < 10; i++ {
		go func() {
			// Each reading will be executed here in range between 1 and 1000 milliseconds.
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			p1.RLock()
			// p1.Name="Trinity"
			// If we try to modify value of this structure under RLock() we will get fatal error.
			// We can only read attributes from the object, we can't change object values.
			// For that we need to use Lock() function for our Mutex.
			fmt.Println(p1.Name)
			p1.RUnlock()
		}()
	}

	go func() {
		time.Sleep(time.Duration(rand.Intn(750)) * time.Millisecond)
		// This Lock will at some point, stop all readings to modify value of Name.
		// After modification, readings will again be allowed.
		// After modification, readings from go routines which are executed in for loop, will continue with reading.
		// But from now, their will print <Agent Smith>, not <Neo> value.
		p1.Lock()
		p1.Name = "Agent Smith"
		p1.Unlock()
	}()
	time.Sleep(1 * time.Second)
}
/*Output:
Neo
Neo
Neo
Neo
Neo
Neo
Agent Smith			> At this point go routine which write new modification, changed value of Name Attribute.
Agent Smith			> After that all other go routines which wanted to read value, continue with their execution.
Agent Smith			> But now, they are reading new value as <Agent Smith>, which is modified by writing Go Routine.
Agent Smith
Agent Smith
 */
