package main

import (
"fmt"
"reflect"
"sync"
)

const mutexLocked = 1

func MutexLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

func RWMutexWriteLocked(rw *sync.RWMutex) bool {
	// RWMutex has a "w" sync.Mutex field for write lock
	state := reflect.ValueOf(rw).Elem().FieldByName("w").FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

func RWMutexReadLockedState(rw *sync.RWMutex) bool {
	return reflect.ValueOf(rw).Elem().FieldByName("readerCount").Int() > 0
}

func RWMutexReadLockedNum(rw *sync.RWMutex) int64 {
	return reflect.ValueOf(rw).Elem().FieldByName("readerCount").Int()
}

func RWMUtexDoubleCheck(rw *sync.RWMutex) bool {
	return reflect.ValueOf(rw).Elem().FieldByName("readerCount").Int() > 0
}

func main() {
	m := sync.Mutex{}
	fmt.Println("m locked =", MutexLocked(&m))	// m locked = false
	m.Lock()
	fmt.Println("m locked =", MutexLocked(&m))	// m locked = true
	m.Unlock()
	fmt.Println("m locked =", MutexLocked(&m))	// m locked = false

	rw := sync.RWMutex{}
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = false  | Readers Num =  0
	rw.Lock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = true  read locked = false  | Readers Num =  -1073741824

	state := reflect.ValueOf(rw).Elem().FieldByName("w").FieldByName("state")
	fmt.Printf("%v > %T",state,state)
	rw.Unlock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = false  | Readers Num =  0
	rw.RLock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = true  | Readers Num =  1
	rw.RLock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = true  | Readers Num =  2
	rw.RUnlock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = true  | Readers Num =  1
	rw.RUnlock()
	fmt.Println("rw write locked =", RWMutexWriteLocked(&rw), " read locked =",
		RWMutexReadLockedState(&rw), " | Readers Num = ", RWMutexReadLockedNum(&rw))
	// rw write locked = false  read locked = false  | Readers Num =  0
}