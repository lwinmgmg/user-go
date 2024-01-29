package maths_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/pkg/maths"
)

func TestSum(t *testing.T) {
	abc := []uint{2, 49, 4}
	res := maths.Sum(abc)
	if res != 55 {
		t.Errorf("Expected 55, getting : %v", res)
	}
}
