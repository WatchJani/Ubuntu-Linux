package main

import (
	"os"
	"root/reader"
	"testing"
)

//bf

const TEST_FILE_PATH string = "./file"

// 500ns per file read / 0 allocation
func BenchmarkOpenRead(b *testing.B) {
	reader := reader.New(100)

	reader.Open(TEST_FILE_PATH)
	for i := 0; i < b.N; i++ {
		reader.Read()
	}

	reader.Close()
}

// 9227ns / 5 allocation
func BenchmarkCloseRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader.Read(TEST_FILE_PATH)
	}
}

// go workers 1000ns / 5 allocation
func BenchmarkWorker(b *testing.B) {
	flow := reader.NewFlow()

	flow.Worker()

	for i := 0; i < b.N; i++ {
		flow.File_name <- "./file"
	}
}

func BenchmarkSmartReader(b *testing.B) {
	reader := reader.NewSmartReader()

	for i := 0; i < b.N; i++ {
		reader.SmartRead()
	}
}

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Open("./file")
	}
}
