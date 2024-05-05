package main


import "fmt"

func main(){
	c,done := make(chan int), make(chan bool)
	go func() {
		for i:=0;i<25;i++{
			c <- i
		}
		close(c)
	}()

	// If we need to read as fast as we can some channel without specific order.
	go func() {
		for n:= range c{
			fmt.Println("Function < 1 >", n)
		}
		done <- true
	}()

	go func() {
		for n:= range c{
			fmt.Println("Function < 2 >", n)
		}
		done <- true
	}()

	for i:=0;i<2;i++{
		<- done
	}
}
/*Output:
Function < 2 > 0
Function < 2 > 1
Function < 2 > 2
Function < 2 > 3
Function < 2 > 4
Function < 1 > 5
Function < 1 > 7
Function < 1 > 8
Function < 1 > 9
Function < 1 > 10
Function < 1 > 11
Function < 1 > 12
Function < 1 > 13
Function < 1 > 14
Function < 1 > 15
Function < 1 > 16
Function < 1 > 17
Function < 1 > 18
Function < 1 > 19
Function < 2 > 6
Function < 2 > 21
Function < 2 > 22
Function < 2 > 23
Function < 2 > 24
Function < 1 > 20
*/

