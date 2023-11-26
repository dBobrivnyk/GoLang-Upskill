package numbers_test

import (
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/numbers"
	"testing"
)

func TestSum(t *testing.T) {
	data := []int{5, 5, 5, 5}
	expected := 20
	res := numbers.Sum(data...)

	if expected != res {
		t.Errorf("Sum == %q, want %q", res, expected)
	}
}

func TestExtendSliceIfSumEven(t *testing.T) {
	data := []int{5, 5, 5, 5}
	expectedLen := 5
	numbers.ExtendSliceIfLenEven(&data, 3)

	if len(data) != expectedLen {
		t.Errorf("Sum == %q, want %q", len(data), expectedLen)
	}
}
