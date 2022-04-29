package tree

import (
	"math/rand"
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
)

func TestPQueue(t *testing.T) {
	num := 25
	p := NewPQueue(num)
	for i := 0; i < num; i++ {
		v := utilcom.RandStringBytesMaskImprSrc(3)
		node := NewPnode(v, rand.Intn(100))
		p.Push(node)
	}
	t.Log(p.IsBigTopHeap())
	for i := 0; i < num; i++ {
		t.Log(p.Pop())
		t.Log(p.IsBigTopHeap())
	}
}

func TestPQueue2(t *testing.T) {
	num := 25
	p := NewPQueue(num)
	for i := 0; i < num; i++ {
		v := utilcom.RandStringBytesMaskImprSrc(3)
		node := NewPnode(v, rand.Intn(100))
		p.Push(node)
		t.Log(p.IsBigTopHeap())
		t.Log(p.Pop())
	}
}
