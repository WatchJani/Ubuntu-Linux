package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("asd")
}

type Worker struct {
	buffer  []byte
	receive chan string
}

func New() *Worker {
	return &Worker{
		receive: make(chan string),
		buffer:  make([]byte, 4096),
	}
}

func Create(num int) {

	for i := 0; i < num; i++ {
		go New()
	}
}

func Worker

func (w Worker) Read() {
	file, _ := os.Open(<-w.receive)
	file.Read(w.buffer)
}
