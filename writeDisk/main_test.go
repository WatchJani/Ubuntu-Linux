package main

import (
	"sync"
	"testing"
)

func Benchmark(b *testing.B) {
	b.StopTimer()
	buffer := make([]byte, 4*1024*1024)
	Full(buffer)

	var wg sync.WaitGroup

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for index := 0; index < len(buffer); index += 4 * 1024 {
			wg.Add(1)
			go Writer(buffer[index:index+4*1024], index, &wg)
		}
		wg.Wait()
	}
}

func BenchmarkWriteInstantInFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recovery("./store/store.bin")
	}
}
