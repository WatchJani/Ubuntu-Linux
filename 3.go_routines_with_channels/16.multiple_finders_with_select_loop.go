package main

import ("time";"math/rand";"fmt")

var scMapping = map[string]int{"David":5,"Jack":10}

func FindSC(name string, c chan int){
	time.Sleep(time.Duration(rand.Intn(6))*time.Second)	// Our random sleep!
	c<-scMapping[name]	// Send to channel value from specific key from map!
}

func main(){
	c1,c2,name1,name2:=make(chan int),make(chan int),"David","Jack"

	go FindSC(name1, c1)	// Sleep will be somewhere between 1 and 6 seconds!
	go FindSC(name2, c2)	// This will also do the same thing!

	go FindSC(name2, c1)
	go FindSC(name2, c2)
	// We will have 4 workers, which will search for the right person for us, who can finish specific job.
	// Select statement allows our code to wait on multiple channels at the same time.
	// In select loop, we can receive only one value from any channel.
	// That means that whoever is first in the line, and whoever find first person suitable for job, will give us value.
	// After that select loop will brake, and we will ignore all other workers, who were searching for that person.
	// We can only receive from one channel in this case.
	// If we put select loop in infinity for loop, then we will expect to get result from multiple workers.
	select {
		case sc := <-c1:
			if sc==5{
				fmt.Println(name1, " with ID ",sc,"found in server_1!")
			} else {
				fmt.Println(name2, " with ID ",sc,"found in server_1!")
			}	
		case sc := <-c2:
			if sc==10{
				fmt.Println(name2, " with ID ",sc,"found in server_2!")
			} else {
				fmt.Println(name1, " with ID ",sc,"found in server_2!")
			}
		case <-time.After(5*time.Second):
		// If no one send the message after 5 seconds, this select block is closed!
			fmt.Println("Search timed out!!!")		
	}
}
// Output: David  with ID  5 found in server_1!
