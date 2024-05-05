package main

import (
	"fmt"
)

func main() {
	dataRead := []string{"data1", "data2", "data3", "data4", "data5", "data6"}

	fmt.Println(dataRead)
}

type Storage struct {
	buffer map[string][]byte
	state  []byte
	stream chan string
}

func New() *Storage {
	return &Storage{
		buffer: make(map[string][]byte, 2),
		state:  make([]byte, 100),
		stream: make(chan string),
	}
}
pol := New()

	pol.WritePolice()
	pol := New()

	pol.WritePolice()
