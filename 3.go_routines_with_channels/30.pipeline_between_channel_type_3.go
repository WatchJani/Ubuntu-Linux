package main

import "fmt"

// In this case our input variable is whole go channel, not specific object.
// Now, as final result, we will return one value of int data type.
func Sum (in <-chan int) int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range in {
			// All our 200 members wil be collected into one int!
			sum+=v
		}
		// REMEMBER: If we don't put this implementation in channel, for <sum> we will get result 0. Why?
		// We started range over <in> channel and summarize all values inside of <sum> variable.
		// BUT, we started that process in isolated go routine.
		// That means that we need to pass value of <sum> variable from that go routine to Sum() function.
		// We will do that with channel, where we will send value of <sum> variable to <out> channel.
		//
		out<-sum
		close(out)
	}()
	// Then here as return we can get value from <out> channel.
	return <-out
}
//==================================================================================
// In this case our input variable is whole go channel, not specific object.
// We will iterate over that channel, so we can process all data, and put new data inside of new channel.
// Then, that new channel we will return as output.
func Power(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for v := range in{	// In this list we have 200 members!
			out<- v*v		// Each member will be powered by himself!
		}
		close(out)
	}()
	return out
}
//==================================================================================
// In this case we return <-chan int.
// That means that we won't return specific object or value with data type.
// We will return whole go channels and all data inside of that channel.
func Generator(max int) <-chan int {
	outChInt := make(chan int)
	// Why we returning whole channel?
	// All buffered values will be sent to next pipeline!
	go func() {
		for i:=1;i<=max;i++{
			outChInt<-i
		}
		close(outChInt)
	}()
	return outChInt
}
//==================================================================================
func LaunchPipeline(amount int) int{
	return Sum(Power(Generator(amount)))
}
//==================================================================================
func main(){
	sum:= LaunchPipeline(200)
	fmt.Println(sum)
}
// Output: 2686700








