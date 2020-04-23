package recursive

import "testing"

func TestFactorial_Recursive(t *testing.T) {
	n := 20
	f := NewFactorial(n)
	t.Log(f.Recursive(n))
}

func TestFactorial_TailRecursive(t *testing.T) {
	n := 13
	f := NewFactorial(n)
	t.Log(f.TailRecursive(n) == f.Recursive(n))
}

func TestFactorial_Loop(t *testing.T) {
	n := 13
	f := NewFactorial(n)
	t.Log((f.Loop(n) == f.Recursive(n)) && (f.Recursive(n) == f.TailRecursive(n)))
}
