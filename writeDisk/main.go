package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func main() {
	// buffer := make([]byte, 8*1024*1024)

	// Full(buffer)

	// var wg sync.WaitGroup

	// wg.Add(8 * 1024 * 1024 / 4096)

	// for index := 0; index < len(buffer); index += 4096 {
	// 	go Writer(buffer[index:index+4096], index, &wg)
	// }

	// wg.Wait()
}

func Full(buffer []byte) {
	for index := range buffer {
		buffer[index] = byte(rand.Intn('Z'-'A'+1) + 'A')
	}
}

func Writer(buffer []byte, name int, wg *sync.WaitGroup) {
	os.WriteFile("./test/"+fmt.Sprintf("%d", name)+".bin", buffer, 0766)

	wg.Done()
}

func Recovery(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	data := []byte("Red 1\n")

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
}
