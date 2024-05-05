package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)
	kickBall := func(playerName string,b chan string) {
		for {
			v, ok := <-b
			if !ok {
				break
			} else if v == "END" {
				close(b)
				break
			} else {
				//This is our trigger for starting our 4 go routines!
				//We block here go routine and only after one seconds we send player name to the channel!
				fmt.Println(v, "kicked the ball.")
				time.Sleep(time.Second)
				b <- playerName
			}
		}
	}
	go kickBall("John",ball)
	go kickBall("Alice",ball)
	go kickBall("Bob",ball)
	go kickBall("Emily",ball)
	//This 4 channel won't start until referee kick the ball.
	ball <- ">>>>> We started game. Referee" // Here we sent new value to ball channel and we started the game.
	e := make(chan bool, 1)
	endTheGame := func(call string,b chan string){
		time.Sleep(5*time.Second)
		b <- call
		e <- true
	}
	endTheGame("END",ball)	// Here we have timer, when we will end the game.
	status := <-e
	if status {
		fmt.Println(">>>>> Match is ended.")
	}
	/*Output:
	>>>>> We started game. Referee kicked the ball.
	Emily kicked the ball.
	Alice kicked the ball.
	Bob kicked the ball.
	John kicked the ball.
	>>>>> Match is ended.
	 */
}
