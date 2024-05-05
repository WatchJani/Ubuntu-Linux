package main

import "fmt"

func increment2(s string) chan int{
	out := make(chan int)
	go func(){
		for i:=0;i<20;i++{
			out <- 1
			fmt.Println(s,i)
		}
		close(out)
	}()
	return out
}

func puller2(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n:= range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func main (){
	c1:=increment2("Incrementor <1> :")
	c2:=increment2("Incrementor <2> :")
	c3:=puller2(c1)
	c4:=puller2(c2)

	fmt.Println("\n 	Final Counter: ", <-c3 + <-c4)
}
/*Output:
Incrementor <2> : 0
Incrementor <2> : 1
Incrementor <2> : 2
Incrementor <2> : 3
Incrementor <2> : 4
Incrementor <2> : 5
Incrementor <2> : 6
Incrementor <2> : 7
Incrementor <2> : 8
Incrementor <2> : 9
Incrementor <2> : 10
Incrementor <2> : 11
Incrementor <2> : 12
Incrementor <2> : 13
Incrementor <2> : 14
Incrementor <2> : 15
Incrementor <2> : 16
Incrementor <2> : 17
Incrementor <2> : 18
Incrementor <2> : 19
Incrementor <1> : 0
Incrementor <1> : 1
Incrementor <1> : 2
Incrementor <1> : 3
Incrementor <1> : 4
Incrementor <1> : 5
Incrementor <1> : 6
Incrementor <1> : 7
Incrementor <1> : 8
Incrementor <1> : 9
Incrementor <1> : 10
Incrementor <1> : 11
Incrementor <1> : 12
Incrementor <1> : 13
Incrementor <1> : 14
Incrementor <1> : 15
Incrementor <1> : 16
Incrementor <1> : 17
Incrementor <1> : 18
Incrementor <1> : 19

 	Final Counter:  40
 */


