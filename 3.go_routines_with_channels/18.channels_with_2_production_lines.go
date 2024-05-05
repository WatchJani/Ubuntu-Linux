package main

import("time";"math/rand";"fmt")

func cakeMaker(kind string, number int, to chan<- string){
	for i:=0; i<number;i++{
		time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
		to<-kind
	}
	close(to)
}

func main(){
	// We need to create specific number of cakes and we have two types of cakes.
	// Two different factories can create both type of cakes.
	// Here we have architecture with multiple production lines and 1 receiver from these productions.
	cake1,cake2 := make(chan string),make(chan string)
	go cakeMaker("CHOCOLATE", 4, cake1)
	go cakeMaker("BANANA", 4, cake2)

	moreCake1,moreCake2,cake := true,true,""
	// This for loop will continue until both channels are closed.
	// If at least one channel is open, we will still be inside of for loop.
	// After both channels are closed or time.After() function is expired, we will brake for loop.
	for moreCake1 || moreCake2 {
		select {
			case cake, moreCake1 = <- cake1:
        //If timeout reach deadline or channel is closed, channel will give us false!
				if moreCake1{
					fmt.Printf("Got a cake from the FIRST factory: %s\n", cake)
				}
			case cake, moreCake2 = <- cake2:
				if moreCake2{
					fmt.Printf("Got a cake from the SECOND factory: %s\n", cake)
				}
			//Everything that is build after 8s won't be sent from factory to us!
			case <-time.After(8000*time.Millisecond):
				fmt.Println("Timed out!")
			return
		}
	}
}

/*One of possible outputs:
Got a cake from the FIRST factory: CHOCOLATE
Got a cake from the SECOND factory: BANANA
Got a cake from the SECOND factory: BANANA
Got a cake from the FIRST factory: CHOCOLATE
Got a cake from the FIRST factory: CHOCOLATE
Got a cake from the SECOND factory: BANANA
Got a cake from the FIRST factory: CHOCOLATE
Got a cake from the SECOND factory: BANANA
 */

	//1 Native Thread > 2MB
	//1 Goroutines > 2kb
	//Resources 1000 Goroutines = Resources 1 Native Thread

	//60% speed up for parsing
	//Speed up of 2x for the actual computation
	//All this adds overhead
	//RAM is slower then CPU
	//Loading from disk is not concurrent
	//ALWAYS MEASURE, DON'T GUESS!