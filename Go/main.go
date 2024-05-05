package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)

	start := time.Now()
	for i := 0; i < 200_000; i++ {
		Read("./AABLNCRYJG.txt", buffer)
	}

	fmt.Println(time.Since(start))
}

func Read(path string, buffer []byte) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Greška pri otvaranju fajla:", err)
		return
	}
	defer file.Close()

	for {
		n, _ := syscall.Read(int(file.Fd()), buffer)

		if n == 0 {
			break
		}
	}
}
