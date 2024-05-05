package main

import (
	"fmt"
	"sync"
)

// Inside of this structure we put map, but also we wrote embedded mutex, so we can get thread safe map.
// This map can be now safely manipulated with multiple go routines.
type RegularIntMap struct {
	sync.RWMutex
	Internal map[string]string
}

// Here we will create constructor for our structure
func NewRegularIntMap() *RegularIntMap {
	return &RegularIntMap{
		Internal: make(map[string]string),
	}
}

// For loading specific key from map, we will use RLock().
// We want to give that type of mutex so some other go routine can also read keys from map at the same time.
func (rm *RegularIntMap) Load(key string) (value string, ok bool) {
	rm.RLock()
	result, ok := rm.Internal[key]
	rm.RUnlock()
	return result, ok
}
// For deleting and storing we will use Lock(), because we want to lock completely that structure in moment of writing.
func (rm *RegularIntMap) Delete(key string) {
	rm.Lock()
	delete(rm.Internal, key)
	rm.Unlock()
}
func (rm *RegularIntMap) Store(key, value string) {
	rm.Lock()
	rm.Internal[key] = value
	rm.Unlock()
}

//==============================================================
func main() {
	sm := NewRegularIntMap()

	sm.Store("Hello", "World") 	// Store an item in the map.
	fmt.Println("Added value: `World` for key: `Hello`")

	result, ok := sm.Load("Hello") 		// Fetch the item we just stored.
	if ok {
		fmt.Printf("Result: `%s` found for key: `Hello`\n", result)
	} else {
		fmt.Println("Value not found for key: `Hello`")
	}

	sm.Delete("Hello")
	sm.Store("Golang", "Go Routines")
	fmt.Println(sm)          // &{{{0 0} 0 0 0 0} map[Golang:Go Routines]}
	fmt.Println(sm.Internal) // map[Golang:Go Routines]
}
