package hashing

import "github.com/google/uuid"

func NewUuid4() string {
	uuid4 := uuid.New()
	return uuid4.String()
}
