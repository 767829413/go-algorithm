package dynamicprogramming

import "testing"

func TestCount(t *testing.T) {
	values := []int{
		100, 600, 700, 99, 49, 129, 229,
	}
	c := NewCount(values, len(values), 350)
	c.count()
}
