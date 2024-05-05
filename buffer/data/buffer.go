package data

type Buffer struct {
	buffer   []byte
	counter  int
	capacity int
}

// make new buffer
func New(capacity int) *Buffer {
	return &Buffer{
		capacity: capacity,
		buffer:   make([]byte, capacity),
	}
}

// append new data to our buffer
func (b *Buffer) Append(data []byte) {
	copy(b.buffer[b.counter:], data)
	b.counter += len(data)
}

// check buffer size with new data
func (b Buffer) Check(data int) bool {
	return b.counter+data > b.capacity
}

// reset buffer
func (b *Buffer) Rewrite() {
	b.counter = 0
}

// return capacity
func (b Buffer) GetCapacity() int {
	return b.capacity
}

// ! delete after testing
func (b *Buffer) Add(data []byte) {
	if b.Check(len(data)) {
		b.Rewrite()
	}

	b.Append(data)
}
