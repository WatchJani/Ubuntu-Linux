package data

import "testing"

// does not depend on the length of the data
func BenchmarkBufferReset(b *testing.B) {
	buffer := New(8_388_608)
	var data []byte = []byte(`{"id": 123,"ime": "John","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}`)

	for i := 0; i < b.N; i++ {
		buffer.Add(data)
	}
}

func BenchmarkAppendData(b *testing.B) {
	buffer := make([]byte, 8_388_608)
	var data []byte = []byte(`{"id": 123,"ime": "John","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}`)

	for i := 0; i < b.N; i++ {
		copy(buffer[0:], data)
	}
}
