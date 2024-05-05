package main

import ("time";"math/rand";"fmt";"sync")

func main() {
	var wg1, wg2 sync.WaitGroup
	// Here we defined two WaitGroups which will wait 6 iterations from 2 different for loops to finish executions.
	// In previous example we had one single WaitGroup with two different checkpoints, and second was waiting first to finish execution.
	// Now we started first checkpoint for first WaitGroup, but now second WaitGroup isn't waiting anymore first checkpoint to finish his execution.
	// Instead of that, second checkpoint is started immediately after first for loop execute his 6 iterations.
	// So, in this case, we have two different checkpoints, where second is working in parallel line with first one.
	// The difference is that we will wait for their finishing line on different place.
	for i := 0;i<5;i++{
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			time.Sleep(time.Duration(rand.Intn(2))*time.Second)
			fmt.Println("WaitGroup - 1 - Work done for ",i)
		}(i)
	}
	for i := 5;i<=10;i++{
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			time.Sleep(time.Duration(rand.Intn(6))*time.Second)
			fmt.Println("WaitGroup - 2 - Work done for ",i)
		}(i)
	}
	//====================================================================
	wg1.Wait()	// Here we will wait all routines from WaitGroup1.
	fmt.Println("First checkpoint finished!")
	fmt.Println("We are sending calculations to data center!")

	wg2.Wait()	// Here we will wait all routines from WaitGroup2.
	fmt.Println("Second checkpoint finished!")

	fmt.Println("All jobs has been finished!")
}
// Here we have two checkpoints working in parallel.
/* Output:
WaitGroup - 1 - Work done for  4
WaitGroup - 2 - Work done for  8
WaitGroup - 2 - Work done for  5
WaitGroup - 1 - Work done for  0
WaitGroup - 1 - Work done for  1
WaitGroup - 2 - Work done for  10
WaitGroup - 1 - Work done for  2
WaitGroup - 1 - Work done for  3
First checkpoint finished!
We are sending calculations to data center!
WaitGroup - 2 - Work done for  9
WaitGroup - 2 - Work done for  6
WaitGroup - 2 - Work done for  7
Second checkpoint finished!
All jobs has been finished!
 */
