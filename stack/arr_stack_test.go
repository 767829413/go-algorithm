package stack

import "testing"

func TestStackChain_Push(t *testing.T) {
	s := NewArrayStack()
	t.Log(s.Print())
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
}

func TestArrayStack_Pop(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
	t.Log(s.Pop())
	t.Log(s.Print())
	t.Log(s.Pop())
	t.Log(s.Print())
}

func TestArrayStack_Top(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
	t.Log(s.top)
}
