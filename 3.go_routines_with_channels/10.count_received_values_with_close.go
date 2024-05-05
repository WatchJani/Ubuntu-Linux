package main

import ("fmt";
	"sync"
	"time")

type Counter struct {
	N 	int
	sync.Mutex
}

func main(){
	ch1,t := make(chan int),&Counter{}

	go addVal(10, ch1, t)
	go addVal(30, ch1, t)
	go addVal(100, ch1, t)

	// We can range over channel because we closed <ch1> channel.
	for i := range ch1 {
		fmt.Println(i, "- Element ",i)
	}
	fmt.Println("We received values from the channel!")
	/* Output:
	Time for execution 230 for 4950!
	4950 - Element  4950
	Time for execution 130 for 45!
	45 - Element  45
	Time for execution 140 for 435!
	435 - Element  435
	We received values from the channel!
	 */
}

func addVal(a int, ch chan int, t *Counter) {
	t.Lock()
	start := time.Now()
	var num int

	for i:=0 ;i<a;i++{
		num=num+i
	}

	elapsed := time.Since(start).Nanoseconds()
	fmt.Printf("Time for execution %v for %v!\n",elapsed,num)

	ch <- num
	t.N++
	if t.N==3{
		close(ch)
	}
	t.Unlock()
}
