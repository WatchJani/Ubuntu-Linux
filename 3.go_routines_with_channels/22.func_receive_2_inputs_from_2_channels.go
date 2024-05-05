package main

import ("time";"math/rand";"fmt")

// Channel r is viewed as a promise by this function.
// Because we put chan<- int32 type, after we send value to that channel, 
// we will immediately read value that we get from the channel in last function!
func LongTimeRequest2(r chan<- int32)  {
	time.Sleep(time.Second * 1)
	r <- rand.Int31n(100)
}

func SumSquares2(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	ra, rb := make(chan int32), make(chan int32)
	go LongTimeRequest2(ra)
	go LongTimeRequest2(rb)

	// This is blocked until both channel received the values and send those values to SumSquares2() function.
	fmt.Println(SumSquares2(<-ra, <-rb))
}
