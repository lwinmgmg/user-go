package hashing

func NewUuid4Hash256() ([]byte, error) {
	uuid4 := NewUuid4()
	return Hash256(uuid4)
}
