package dynamicprogramming

import "testing"

func TestGetMinDest1(t *testing.T) {
	w := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	d := NewDPStatusTable(w, len(w))
	t.Log(d.GetMinDest())
}

func TestGetMinDest2(t *testing.T) {
	w := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	d := NewDPStatusFunc(w, len(w))
	t.Log(d.GetMinDest())
}
