package queue

import (
	"sync"
	"testing"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	n := 5
	q := NewCircularQueue(n)
	for i := 0; i < n; i++ {
		q.EnQueue(i)
	}
	t.Log(q.Print())
}

func TestCircularQueue_DeQueue(t *testing.T) {
	n := 5
	q := NewCircularQueue(n)
	for i := 0; i < n; i++ {
		q.EnQueue(i)
	}
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
	t.Log(q.DeQueue())
	q.EnQueue(22)
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
	t.Log(q.DeQueue())
	t.Log(q.Print())
	q.EnQueue(33)
	q.EnQueue(36)
	q.EnQueue(37)
	q.EnQueue(38)
	t.Log(q.Print())
}

func TestCircularQueue_produce_consume(t *testing.T) {
	n := 6
	q := NewCircularQueue(n)
	var g sync.WaitGroup
	for i := 0; i < n; i++ {
		g.Add(2)
		go func(i int) {
			defer g.Done()
			t.Log("插入", i, q.tail, q.head, q.EnQueue(i))
		}(i)

		go func() {
			defer g.Done()
			t.Log("拿出", q.tail, q.head, q.DeQueue())
		}()
	}
	g.Wait()
}

func Test_produce_consume_channel(t *testing.T) {
	ch := make(chan interface{}, 6)
	var g sync.WaitGroup
	for i := 0; i < 20; i++ {
		g.Add(2)
		go func(i int) {
			defer g.Done()
			ch <- i
			t.Log("插入", i)
		}(i)

		go func() {
			defer g.Done()
			t.Log("拿出", <-ch)
		}()
	}
	g.Wait()
}
