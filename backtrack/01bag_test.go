package backtrack

import "testing"

func TestBagQuestion(t *testing.T) {
	items := []int{
		5, 1, 9, 4,
	}
	BagQuestion(0, 0, items, len(items), 12)
	t.Log(GetMax())
}
