package main

import (
	"fmt"
	"time"
	"sync"
)

type MinTime struct {
	Elapsed		int64
	Value 		int
	sync.RWMutex
}
var minTime MinTime
var t = 0

func add(a int, ch chan map[int64]int, size int) {
	start := time.Now()
	var num int

	for i:=0 ;i<a;i++{
		num=num+i
	}
	elapsed := time.Since(start).Nanoseconds()
	fmt.Printf("Time for execution %v for %v!\n",elapsed,num)

	m1 := make(map[int64]int)
	m1[elapsed]=num

	ch <- m1
	t++
	if t==size{
		close(ch)
		t=0
	}
}

func (mT *MinTime) receiveMap (ch1 chan map[int64]int,g1 *sync.WaitGroup) {
	m1 := <-ch1
	for k, v := range m1 {
		if mT.Elapsed == 0.0 {
			mT.Elapsed = k
		} else {
			if mT.Elapsed > k {
				mT.Lock()
				mT.Elapsed = k
				mT.Value = v
				mT.Unlock()
			}
		}
	}
	g1.Done()
}

func main() {
	ch1 := make(chan map[int64]int)
	slice1 := []int{10, 30, 50, 70, 100, 120, 150, 200, 210, 240, 250, 270, 290, 300}

	for i := 0; i < len(slice1); i++ {
		go add(slice1[i], ch1, len(slice1))
	}
	var wg sync.WaitGroup
	wg.Add(len(slice1))
	for i:=0;i<len(slice1);i++{
		go minTime.receiveMap(ch1,&wg)
	}
	wg.Wait()
	fmt.Println("Result for minTime was", minTime.Value, "! minTime was", minTime.Elapsed, "nanoseconds!")
	fmt.Println("We received values from the channel!")
}
/*Output:
Time for execution 300 for 45!
Time for execution 160 for 2415!
Time for execution 120 for 435!
Time for execution 330 for 44850!
Time for execution 310 for 7140!
Time for execution 230 for 11175!
Time for execution 350 for 19900!
Time for execution 280 for 1225!
Time for execution 320 for 21945!
Time for execution 270 for 28680!
Time for execution 490 for 41905!
Time for execution 270 for 4950!
Time for execution 160 for 36315!
Time for execution 340 for 31125!
Result for minTime was 435 ! minTime was 120 nanoseconds!
We received values from the channel!
 */




