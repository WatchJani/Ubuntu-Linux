package reader

type SmartReader struct {
	buffer []byte
}

func NewSmartReader() *SmartReader {
	return &SmartReader{
		buffer: make([]byte, 4096),
	}
}

func (s *SmartReader) SmartRead() []byte {
	return s.buffer[200:258]
}
