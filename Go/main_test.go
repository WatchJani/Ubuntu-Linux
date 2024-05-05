package main

import (
	"os"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	b.StopTimer()
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Read("./AABLNCRYJG.txt", buffer)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.ReadFile("./AABLNCRYJG.txt")
	}
}

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Open("./AABLNCRYJG.txt")
	}
}

func BenchmarkOpen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.OpenFile("./AABLNCRYJG.txt", os.O_RDONLY, 0)
	}
}
