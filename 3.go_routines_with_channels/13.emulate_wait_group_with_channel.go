package main

import "fmt"

func Sum2(s []int,c chan int){
	sum:=0
	for _,v := range s{
		sum+=v
	}
	c<-sum
}

func main(){
	s,c := []int{8,4,3,7,-4,-2,-5},make(chan int)

	fmt.Println(s[:len(s)/2])	// First 3 elements: [8 4 3]
	fmt.Println(s[len(s)/2:])	// We start after 3rd element: [7 -4 -2 -5]

	go Sum2(s[:len(s)/2],c)		// 15
	go Sum2(s[len(s)/2:],c)		// -4

	// Here we emulated WaitGroup without WaitGroup.
	// Values for <x> and <y> variables will be blocking mechanism, until we receive both values from functions.
	x,y := <-c,<-c
	fmt.Println(x,y,x+y)		// -4, 15, 11
}
