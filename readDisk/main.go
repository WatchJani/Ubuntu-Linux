package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	Read("16.bin")
}

func Read(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	newReader := bufio.NewScanner(file)

	for newReader.Scan() {

	}
}

func ReadAll(path string) []byte {
	buf, _ := os.ReadFile(path)

	return buf
}

func ReadByte(path string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			//parser
			if string(buffer[:n]) == "da" {
				fmt.Println("yes")
			}
			// rukovanje greškom
		}
		// radite sa podacima u baferu (prvih n bajtova)
	}
	wg.Done()
}
