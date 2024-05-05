package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var counter1 int32
var counter2 int64
//var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Here both functions <Exec1> and <Exec2> are increasing values of global variables.
	go Exec1(&wg)
	go Exec2(&wg)

	// WaitGroup is waiting both functions to finish their executions.
	wg.Wait()
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))	// Here value is 32, because I have 16 cores and 32 threads.
	// In my case this is 8 virtual thread cores, and 4 physical cores
	fmt.Println("Final Counter: ", counter1)
}

func Exec1(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// Here we can write mutex or we can put variables in Atomic Functions
		//mutex.Lock()
		// You can try to put incrementing with <++> operator, without Atomic Functions, so you can see the difference.
		// counter1++
		// counter2++
		atomic.AddInt32(&counter1, 2)
		atomic.AddInt64(&counter2, 2)
		fmt.Println("Exec<1>: ", "Iteration =",i, "Counter <1> Value =",counter1)
		fmt.Println("Exec<1>: ", "Iteration =",i, "Counter <2> Value =",counter2)
		//mutex.Unlock()
	}
	fmt.Println("=============Execution of go routine num <1> is finished!=============")
	wg.Done()
}

func Exec2(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// Here we can write mutex or we can put variables in Atomic Functions
		//mutex.Lock()
		// You can try to put incrementing with <++> operator, without Atomic Functions, so you can see the difference.
		// counter1++
		// counter2++
		atomic.AddInt32(&counter1, 3)
		atomic.AddInt64(&counter2, 3)
		fmt.Println("Exec<2>: ", "Iteration =",i, "Counter <1> Value =",counter1)
		fmt.Println("Exec<2>: ", "Iteration =",i, "Counter <2> Value =",counter2)
		//mutex.Unlock()
	}
	fmt.Println("=============Execution of go routine num <2> is finished!=============")
	wg.Done()
}

// CONCURRENCY is the composition of independently executing processes
// PARALLELISM is the simultaneous execution of(possibly related) computations.

// CONCURRENCY is about dealing with lots of things at once.
// PARALLELISM is about doing lots of things at once.
//==============================================================================
// In CONCURRENCY I can drink tea and then talk to my student.
// In PARALLELISM it's "impossible" to drink and talk to my student.

// In PARALLELISM i can write on my keyboard, and watch the second monitor.
// CONCURRENCY will run on one core by default!
// If you want more, tell him specifically that you want to use more then one!

/*
CONCURRENCY								PARALLELISM >> MULTI-CORE JOBS
*****************************			************************************
							*			*		A
A					A		*			*	----------------------->
---->               ------>	*			*		B
               B			*	  VS	*	----------------------->
            --->			*			*		C
	 C             C		*			*	----------------------->
	 ------>    -->			*			*
							*			*
*****************************			************************************

*/
/*
Foo:  0 2
Foo:  1 4
Foo:  2 6
Foo:  3 8
Foo:  4 10
Foo:  5 12
Bar:  0 14
Foo:  6 16
Foo:  7 18
Foo:  8 20
Foo:  9 22
Bar:  1 24
Bar:  2 26
Bar:  3 28
Bar:  4 30
Bar:  5 32
Bar:  6 34
Bar:  7 36
Bar:  8 38
Bar:  9 40
*/
