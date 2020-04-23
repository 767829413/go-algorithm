package queue

import "testing"

func TestLinkedListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	t.Log(q.Print())
}

func TestLinkedListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	for i := 0; i < 3; i++ {
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
