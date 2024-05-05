package main

import (
	"sync"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read("4.bin")
	}
}

func BenchmarkReadAll(b *testing.B) {
	b.StopTimer()
	var buf []byte
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf = ReadAll("4.bin")
	}

	_ = buf
}

func BenchmarkReadByte(b *testing.B) {
	b.StopTimer()

	var wg sync.WaitGroup

	b.StartTimer()

	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go ReadByte("4.bin", &wg)
	}

	wg.Wait()
}
