package main

import ("sync";"fmt";"time";"math/rand")

type Data struct {
	data []int
	readCount int64
	mutex sync.Mutex
}

// Here multiple writers will operate over one structure, with two different functions.
// The difference is that first function Writer() will be connected with WaitGroup.
// After WaitGroup is finished, we will exit from our scope.
func (d *Data) Writer(g1 *sync.WaitGroup) {
	for i:=1; i<6 ;i++{
		// We will lock
		d.mutex.Lock()
		fmt.Printf("*******> : PERFORMING WRITE : Reading Go Routine Number [%d] - No one is reading at the moment *******\n", d.readCount)
		d.data = append(d.data,i)
		d.mutex.Unlock()
		g1.Done()
		time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
	}
}
// This is still type of writer because we we changed value of <readCount> variable.
func (d *Data) ReadAndUpdate(id int) {
	for{
		d.mutex.Lock()
		d.readCount++
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		fmt.Printf("Reading Go Routine Number [%d] > : Performing Read and Update: Length[%d] Count[%d]\n", id, len(d.data), d.readCount)
		d.readCount--
		d.mutex.Unlock()
	}
}
//======================================================================================================================
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	data := &Data{}
	go data.Writer(&wg)

	// Here we setup 6 different readers, to read values from data and counter.
	for i:=1; i<7 ;i++{
		go data.ReadAndUpdate(i)
	}
	// After 5 writings, WaitGroup will be finished.
	// After that we will move on, and we will exit out of scope.
	wg.Wait()
	fmt.Println("All writes finished their work!")
}
/*
Reading Go Routine Number [6] > : Performing Read and Update: Length[0] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[0] Count[1]
*******> : PERFORMING WRITE : Reading Go Routine Number [0] - No one is reading at the moment *******
Reading Go Routine Number [4] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[1] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[1] Count[1]
*******> : PERFORMING WRITE : Reading Go Routine Number [0] - No one is reading at the moment *******
Reading Go Routine Number [5] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[2] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[2] Count[1]
*******> : PERFORMING WRITE : Reading Go Routine Number [0] - No one is reading at the moment *******
Reading Go Routine Number [2] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[3] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[3] Count[1]
*******> : PERFORMING WRITE : Reading Go Routine Number [0] - No one is reading at the moment *******
Reading Go Routine Number [5] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [4] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [5] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [1] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [3] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [2] > : Performing Read and Update: Length[4] Count[1]
Reading Go Routine Number [6] > : Performing Read and Update: Length[4] Count[1]
*******> : PERFORMING WRITE : Reading Go Routine Number [0] - No one is reading at the moment *******
Reading Go Routine Number [4] > : Performing Read and Update: Length[5] Count[1]
All writes finished their work!


 */