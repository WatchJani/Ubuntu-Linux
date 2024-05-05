package main

import "fmt"


func syncRead(chs ...chan int) chan []int {
	outChan := make(chan []int, 16)
	go func() {
		defer close(outChan)
		for rs, ok := recvOneEach(chs...); ok; rs, ok = recvOneEach(chs...) {
			outChan <- rs
		}
	}()
	return outChan
}

func recvOneEach(chs ...chan int) (rs []int, ok bool) {
	ok = true
	for _, ch := range chs {
		r, ok2 := <-ch
		rs, ok = append(rs, r), ok && ok2
	}
	return rs, ok
}

func main() {
	bufSize := 1
	ch1, ch2, ch3 := make(chan int, bufSize), make(chan int, bufSize), make(chan int, bufSize)
	go func() {
		defer close(ch1)
		defer close(ch2)
		defer close(ch3)
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
		ch3 <- 6
		ch2 <- 5
		ch1 <- 4
	}()
	for rs := range syncRead(ch1, ch2, ch3) {
		fmt.Println(rs)
	}
}