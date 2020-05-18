package tree

import (
	"testing"
)

func TestFullRange(t *testing.T) {
	arr := []int{
		2, 5, 8, 9, 11, 12,
	}
	num := 0
	FullRange(arr, len(arr), len(arr), &num)
	t.Log(num)
}

func TestCellBreed(t *testing.T) {
	t.Log(CellBreed(5))
}
