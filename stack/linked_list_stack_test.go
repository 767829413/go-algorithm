package stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	s := NewLinkedListStack()
	t.Log(s.Print())
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
}

func TestLinkedListStack_Pop(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
	t.Log(s.Pop())
	t.Log(s.Print())
	t.Log(s.Pop())
	t.Log(s.Print())
}

func TestLinkedListStack_Top(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	t.Log(s.Print())
	t.Log(s.Top())
}
