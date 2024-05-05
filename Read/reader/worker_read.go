package reader

import "os"

const NUMBER_OF_WORKERS int = 20

type Flow struct {
	File_name chan string
}

func NewFlow() *Flow {
	return &Flow{
		File_name: make(chan string),
	}
}

func (f *Flow) Worker() {
	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		go f.Read()
	}
}

func (f *Flow) Read() {
	for name := range f.File_name {
		os.ReadFile(name)
	}
}

func (f *Flow) Send() {
	f.File_name <- "./file"
}
