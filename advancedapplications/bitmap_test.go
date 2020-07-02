package advancedapplications

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	num := 20
	b := NewBitMap(num)
	b.Set(2)
	b.Set(4)
	b.Set(6)
	b.Set(8)
	b.Set(10)
	b.Set(12)
	b.Set(num)
	for i := 0; i < num+1; i++ {
		t.Log(i, b.Get(i))
	}
	t.Log(b.bytes)
}

func TestTTTT(t *testing.T) {
	fmt.Println(100>>3 + 1)
}
