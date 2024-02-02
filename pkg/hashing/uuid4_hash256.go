package hashing

func NewUuid4Hash256() ([]byte, error) {
	uuid4 := NewUuid4()
	return Hash256(uuid4)
}

func NewUuid4Hash256Hex() (string, error) {
	uuid4 := NewUuid4()
	return Hash256Hex(uuid4)
}
