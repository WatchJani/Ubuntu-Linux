package main

import (
	"fmt"
	"time"
	"math/rand"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	time.Sleep(time.Duration(rb) * time.Second) // sleep 1s, 2s or 3s
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	c := make(chan int32, 5) // need a buffered channel
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <- c // only the first response is used
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
