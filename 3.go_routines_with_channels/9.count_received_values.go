package main

import ("fmt")

func Add(a int, ch chan int) {
	var num int
	for i:=0 ;i<a;i++{
		num=num+i
	}
	ch <- num
}

func main(){
	ch1,ch2,t := make(chan int),make(chan int),0

	go Add(10, ch1)
	go Add(100, ch2)

	for {
		select {
			case msg := <-ch1:
				fmt.Printf("#CH1 returned %d\n", msg)
			case msg := <-ch2:
				fmt.Printf("#CH2 returned %d\n", msg)
		}
		t++
		// When we receive second value from channels, for loop will brake.
		if t==2 {break}
	}
	fmt.Println("All channels has finished their work!")
	// Output:
	// #CH2 returned 4950
	// #CH1 returned 45
	// All channels has finished their work!
}