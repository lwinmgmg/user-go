package hashing_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/pkg/hashing"
)

func TestHash256(t *testing.T) {
	data, err := hashing.Hash256("LMM")
	if data == nil || err != nil {
		t.Errorf("Error on hashing 256 : %v, %v", data, err)
	}
}
