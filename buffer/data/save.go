package data

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"path"
)

// buffer default size
const MB_8 int = 8_388_608

type Save struct {
	*Buffer
	stream chan []byte
}

// Create new save
func NewSave() *Save {
	return &Save{
		Buffer: New(MB_8),
		stream: make(chan []byte),
	}
}

// add new element, if buffer is full then write on disk
func (s *Save) Save(data []byte) {
	if s.Check(len(data)) {
		s.stream <- s.buffer[:s.counter]
		s.Rewrite()
	}

	s.Append(data)
}

// Close input sender
func (s *Save) Close() {
	close(s.stream)
}

// single writer on disk
func (s Save) Listen() {
	go func() {
		for data := range s.stream {
			if err := Write(data); err != nil {
				log.Println(err)
			}
		}
	}()
}

// write all from buffer in new generated file
func Write(data []byte) error {
	name, err := GenerateRandomString(10)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join("test", name), data, 0644)
}

// generate random string from file name (for now is good)
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
