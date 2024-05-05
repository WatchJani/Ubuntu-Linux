package main

import ("fmt";"sync")

func main() {
	// Previous example can be also defined like this.
	// We can call <sync.Map> which already have pre-defined functions, like we had in previous example.
	var sm sync.Map								// Create the thread safe map.
	sm.Store("Hello", "World")		// Store an item in the map.
	fmt.Println("Added value: `World` for key: `Hello`")

	result, ok := sm.Load("Hello")	// Fetch the item we just stored.
	if ok {
		fmt.Println(result)				// World
	} else {
		fmt.Println("Value not found for key: `Hello`")
	}

	sm.Delete("Hello")
	fmt.Println(sm)
}
