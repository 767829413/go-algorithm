package tree

import (
	"algorithm/utilcom/test"
	"testing"
)

func TestHeap(t *testing.T) {
	num := 120
	arr := test.BuildTestArr(num, 100, 0)
	h := NewHeap(num)
	for _, v := range arr {
		h.Insert(v)
	}
	t.Log(arr)
	t.Log("heap")
	t.Log(h.v)
	t.Log(IsBigTopHeap(h.v, h.count))
	n := h.count
	for i := 0; i < n; i++ {
		h.DelMax()
		t.Log(h.v)
		t.Log(IsBigTopHeap(h.v, h.count))
	}
}
