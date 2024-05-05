package main

import (
	"fmt"
	"log"
	"time"
)

var scMapping2 = map[string]int{"David": 5, "Jack": 10}

func FindSC2(name string, c chan int) {
	time.Sleep(4 * time.Second) // Our random sleep!
	c <- scMapping2[name]       // Send to channel value from specific key from map!
}

func main() {
	t := time.Now()
	c1, c2, name1 := make(chan int), make(chan int), "David"

	go FindSC2(name1, c1) // Sleep will be somewhere between 1 and 4 seconds.
	go FindSC2(name1, c2) // This will also do the same thing.

	// We can only receive from one channel in this case. We need infinity for loop.
	// Select statement allows our code to wait on multiple channels at the same time.
	// Select blocks until one channel is ready to receive the value.
	b := false
	for {
		select {
		case sc := <-c1:
			fmt.Println(name1, " with ID ", sc, "found in server_1!")
		case sc := <-c2:
			fmt.Println(name1, " with ID ", sc, "found in server_2!")
		case <-time.After(5 * time.Second):
			//If no one send the message after 5 seconds, this select block is closed!
			fmt.Println("Search timed out!!!")
			b=true
		}
		if b {
			break
		}
	}

	// Possible Output:
	// 						David  with ID  5 found in server_1!
	// 						David  with ID  5 found in server_2!

	elapsed := time.Since(t)
	log.Printf("Binomial took %s", elapsed)		// 9
}