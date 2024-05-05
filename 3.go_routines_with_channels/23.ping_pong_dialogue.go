package main

import (
	"fmt"
	"os"
	"time"
)

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball,ok := <-table // get the ball
		if !ok {
			break
		} else if ball == 0 {
			close(table)
			break
		} else {
			fmt.Println(playerName, ball)
			ball += lastValue
			if ball < lastValue { 	// overflow
				os.Exit(0)
			}
			lastValue = ball
			time.Sleep(time.Second)
			table <- ball 			// bat back the ball
		}
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		table <- 1 // Throw ball on table
	}()
	// Ping-Pong mechanism we can share data between multiple go routines and stop them whenever we want.
	go Play("A:", table)
	go Play("B:", table)

	e := make(chan bool, 1)
	endTheGame := func(call Ball,b chan Ball){
		time.Sleep(4*time.Second)
		b <- call
		e <- true
	}
	endTheGame(0,table)	// Here we have timer, when we will end the game.
	status := <-e
	if status {
		fmt.Println(">>>>> Match is ended.")
	}
}
