package backtrack

import "testing"

func TestMinDistBT(t *testing.T) {
	w := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	t.Log(GetMinDest(0, 0, 0, len(w), w))
}
