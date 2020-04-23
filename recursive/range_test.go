package recursive

import "testing"

func TestRangeType_RangeALL(t *testing.T) {
	slice1 := NewRangeType(4)
	for i := 0; i < 4; i++ {
		slice1.data[i] = i + 1
	}
	slice1.RangeALL(0)

	slice2 := NewRangeType(3)
	slice2.data[0] = "a"
	slice2.data[1] = "b"
	slice2.data[2] = "c"
	slice2.RangeALL(0)
}
