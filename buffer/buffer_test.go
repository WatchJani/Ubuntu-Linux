package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"root/data"
	"testing"
)

func BenchmarkBuffer(b *testing.B) {
	var buffer []byte = make([]byte, 0, 4096)

	textToAdd := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	for i := 0; i < b.N; i++ {
		buffer = append(buffer, textToAdd...)
	}
}

func BenchmarkBufferDefault(b *testing.B) {
	var buffer bytes.Buffer
	buffer.Grow(4096)

	textToAdd := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	for i := 0; i < b.N; i++ {
		buffer.Write(textToAdd)
	}
}

func BenchmarkBufferMy(b *testing.B) {
	var buffer []byte = make([]byte, 4096)

	textToAdd := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	for i := 0; i < b.N; i++ {
		copy(buffer[1020:], textToAdd)
	}
}

func BenchmarkWriteOnDisck(b *testing.B) {
	client3 := []byte(`Janko", "Kondić", "JankoKondic2722", "jankokondic84@gmail.com", "+386 66 311 063`)
	save := data.NewSave()

	save.Listen()

	for i := 0; i < b.N; i++ {
		save.Save(client3)
	}

	save.Close() //close chanel, for performance test
}

// ima ne greska
func BenchmarkWrite(b *testing.B) {
	client, err := os.ReadFile("test/0b5a13a708af3e641aa0")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < b.N; i++ {
		name, err := data.GenerateRandomString(10)
		if err != nil {
			fmt.Println(err)
		}
		os.WriteFile(path.Join("test", name), client, 0644)
	}
}
