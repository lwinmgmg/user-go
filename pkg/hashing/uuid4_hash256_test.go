package hashing_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/pkg/hashing"
)

func TestUuid4Hash256(t *testing.T) {
	newUuid, err := hashing.NewUuid4Hash256()
	if err != nil {
		t.Errorf("Error on getting new uuid4hash256 : %v", err)
	}
	if len(newUuid) == 0 {
		t.Error("No uuid4 is generated")
	}
	newUuid1, err := hashing.NewUuid4Hash256Hex()
	if err != nil {
		t.Errorf("Error on getting new uuid4hash256 : %v", err)
	}
	if len(newUuid1) == 0 {
		t.Error("No uuid4 is generated")
	}
}
