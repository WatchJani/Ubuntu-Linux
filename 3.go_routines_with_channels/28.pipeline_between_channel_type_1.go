package main

import "fmt"

func increment() chan int{
	out := make(chan int)
	go func(){
		for i:=1;i<10;i++{
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	var sum int
	//After we finish range over input channel, we sent value on channel, close channel, and return chan int type!
	go func() {
		for n:= range c {
			fmt.Println(n)
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func main (){
	c:=increment()
	cSum:=puller(c)
	for n:= range cSum{
		fmt.Println(n)
	}
}
/*Output:
1
2
3
4
5
6
7
8
9
45
 */


