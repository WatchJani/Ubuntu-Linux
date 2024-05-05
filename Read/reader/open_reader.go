package reader

import (
	"fmt"
	"io"
	"os"
)

type Reader struct {
	buffer []byte
	file   *os.File
}

func New(capacity int) *Reader {
	return &Reader{
		buffer: make([]byte, capacity),
	}
}

func (r *Reader) Open(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	r.file = file
}

func (r *Reader) Close() {
	r.file.Close()
}

func (r *Reader) Read() {
	if r.file == nil {
		fmt.Println("file nije otvoren!")
	}

	io.ReadFull(r.file, r.buffer)
}
