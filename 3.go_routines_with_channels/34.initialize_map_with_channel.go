package main

import "fmt"

func main() {

	things := make(map[string]chan int)

	things["stuff"] = make(chan int, 2)
	things["stuff"] <- 2

	things["item"] = make(chan int, 2)
	things["item"] <- 3

	mything := <-things["stuff"]
	mything2 := <-things["item"]

	fmt.Printf("my thing: %d\n", mything)
	fmt.Printf("my thing2: %d\n", mything2)
}
