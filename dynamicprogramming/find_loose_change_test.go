package dynamicprogramming

import "testing"

func TestFindLooseChange(t *testing.T) {
	changeType := []int{
		1, 4, 3, 5,
	}
	f := NewFindLooseChange(changeType, len(changeType), 9)
	t.Log(f.GetMin())
	t.Log(f.GetMin2())
}
