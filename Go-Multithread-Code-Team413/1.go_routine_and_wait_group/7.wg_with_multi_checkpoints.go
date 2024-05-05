package main

import ("time";"math/rand";"fmt";"sync")

func main() {
	// Here we defined one single wait group, but we have two different checkpoints.
	// After we finish first checkpoint, we can start to create new one by adding to number for waiting processes.
	// In this case, we have 6 iterations in first, and 6 iterations with second checkpoint.
	// Each checkpoint will be finished in 3 seconds and each second we will have specific number of finished processes.
	var wg sync.WaitGroup
	for i := 0;i<=5;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3))*time.Second)
			fmt.Println("Work done for ",i)
		}(i)
	}
	wg.Wait()
	fmt.Println("First checkpoint finished!")
	for i := 5;i<=10;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
	// We have 3 checkpoints: after 1,2 and 3 seconds. 
	// Maybe in first we have 2, in second we have 1, and in the last we have 3!
			time.Sleep(time.Duration(rand.Intn(3))*time.Second)
			fmt.Println("Work done for ",i)
		}(i)
	}
	wg.Wait()
	fmt.Println("All jobs has been finished!")
}
