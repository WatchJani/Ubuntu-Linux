package main

import (
	"fmt"
	"time"
)

func worker413(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 1; w <= 3; w++ {
		go worker413(w, jobs, results)
	}

	/*This will give us an error, because we didn't close <results> channel:
	for v := range results {
		fmt.Println(v)
	}*/
	for a := 1; a <= 5; a++ {
		<-results
	}
	/* One of possible outputs:
	worker 3 started  job 1
	worker 1 started  job 2
	worker 2 started  job 3
	worker 3 finished job 1
	worker 3 started  job 4
	worker 2 finished job 3
	worker 2 started  job 5
	worker 1 finished job 2
	worker 2 finished job 5
	worker 3 finished job 4
	*/
}
