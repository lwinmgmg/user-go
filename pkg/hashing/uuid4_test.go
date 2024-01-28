package hashing_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/pkg/hashing"
)

func TestUuid4(t *testing.T) {
	newUuid := hashing.NewUuid4()
	newUuid1 := hashing.NewUuid4()
	if newUuid == "" || newUuid1 == "" {
		t.Errorf("Getting Empty uuid4")
	}
	if newUuid == newUuid1 {
		t.Errorf("Uuid4 must be unique : %v : %v", newUuid, newUuid1)
	}
}
