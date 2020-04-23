package recursive

import "testing"

func TestFibonacci_Recursive(t *testing.T) {
	n := 40
	f := NewFibonacci(n)
	t.Log(f.Recursive(n))
}

func TestFibonacci_TailRecursive(t *testing.T) {
	n := 400
	f := NewFibonacci(n)
	t.Log(f.TailRecursive(n))
}

func TestFibonacci_Loop(t *testing.T) {
	n := 400
	f := NewFibonacci(n)
	t.Log(f.Loop(n))
}
