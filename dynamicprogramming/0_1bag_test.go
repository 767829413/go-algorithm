package dynamicprogramming

import "testing"

func TestBageg1_GetMax(t *testing.T) {
	lw := 20
	items := []int{
		7, 6, 3, 12,
	}
	b := Newbageg1(lw, len(items), items)
	t.Log(b.GetMax(0, 0))
}

func TestBageg2_GetMax(t *testing.T) {
	lw := 20
	items := []int{
		7, 6, 3, 12,
	}
	b := Newbageg2(len(items), lw, items)
	t.Log(b.GetMax())
}

func TestBageg3_GetMax(t *testing.T) {
	lw := 20
	items := []int{
		7, 6, 3, 12,
	}
	b := Newbageg3(len(items), lw, items)
	t.Log(b.GetMax())
}

func TestBageg4_GetMax(t *testing.T) {
	lw := 20
	items := []int{
		7, 6, 3, 12,
	}
	values := []int{
		33, 50, 18, 100,
	}
	b := Newbageg4(len(items), lw, items, values)
	t.Log(b.GetMax())
}
