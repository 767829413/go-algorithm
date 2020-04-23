package queue

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	n := 5
	q := NewArrayQueue(n)
	for i := 0; i < n; i++ {
		q.EnQueue(i)
	}
	t.Log(q.Print())
	t.Log(q.EnQueue(12))
}

func TestArrayQueue_DeQueue(t *testing.T) {
	n := 5
	q := NewArrayQueue(n)
	for i := 0; i < n; i++ {
		q.EnQueue(i)
	}
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
}
