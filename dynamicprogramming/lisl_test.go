package dynamicprogramming

import "testing"

func TestLISL_GetMaxLength(t *testing.T) {
	arr := []int{2, 9, 3, 6, 5, 1, 7,}
	//arr := []int{10, 9, 2, 5, 3, 7, 101, 18,}
	//arr := []int{80, 90, 100, 110, 20, 21, 22, 23, 25, 27, 1, 2, 3, 4, 5,}
	l := NewLISL(arr, len(arr))
	t.Log(l.GetMaxLength())
}
