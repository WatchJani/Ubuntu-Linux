package main

import ("time";"fmt")

func main() {
	checkpoint1 := time.Tick(100 * time.Millisecond)
	checkpoint2 := time.Tick(500 * time.Millisecond)
	checkpoint3 := time.After(900 * time.Millisecond)
	endOfCycle := time.After(2000 * time.Millisecond)

	// With select loop inside of for infinity loop, we can create multiple checkpoints with specific time interval.
	for {
		select {
		case <-checkpoint1:
			// This code will be executed every 100ms as Checkpoint 1.
			fmt.Println("We took data from Data Center 1.")
		case <-checkpoint2:
			// This code will be executed every 100ms as Checkpoint 1.
			fmt.Println("We took data from Data Center 2.")
		case <-checkpoint3:
			// This code will be executed only once after 900ms as Checkpoint 3.
			// It will be executed only once, because we defined checkpoint with After() function, not with Tick().
			fmt.Println("We took ALL DATA from Data Center 3 with ONE pulling.")
		case <-endOfCycle:
			// After 1050ms we will end this cycle.
			// Here we break the infinity loop.
			fmt.Println("Cycle of pulling data is finished.")
			return
		// In default setting we can setup processing which will be executed between checkpoints at specific moment.
		default:
			fmt.Println("Waiting for next checkpoint...")			  // This will start first!
			time.Sleep(200 * time.Millisecond) // In this case, deafult will start first!
		}
	}
}
/* Output:
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 1.
We took data from Data Center 2.
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 2.
We took ALL DATA from Data Center 3 with ONE pulling.
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 2.
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 1.
Waiting for next checkpoint...
We took data from Data Center 2.
Cycle of pulling data is finished.
 */