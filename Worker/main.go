package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUMBER_OF_WORKERS_READER int = 50

func main() {
	stream := Simulation()

	var wg sync.WaitGroup
	wg.Add(NUMBER_OF_WORKERS_READER)
	for i := 0; i < NUMBER_OF_WORKERS_READER; i++ {
		go Read(i, stream, &wg)
	}

	wg.Wait()
}

func Simulation() <-chan struct{} {
	streaming := make(chan struct{})

	go func() {
		for i := 0; i < 1000; i++ {
			streaming <- struct{}{}
		}
		close(streaming)
	}()

	return streaming
}

func Read(id int, stream <-chan struct{}, wg *sync.WaitGroup) {
	for range stream {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println(id)
	}

	wg.Done()
}
