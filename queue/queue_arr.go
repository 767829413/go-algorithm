package queue

import (
	"sync"
)

type Queue struct {
	Mutx  *sync.RWMutex
	Items []string
	Count int
}

func NewQueue(n int) *Queue {
	return &Queue{
		Mutx:  &sync.RWMutex{},
		Items: []string{},
		Count: n,
	}
}

//入队
func (q *Queue) EnQueue(item string) bool {
	if len(q.Items) == q.Count {
		return false
	}
	q.Items = append(q.Items, item)
	return true
}

//出队
func (q *Queue) DeQueue() (item string) {
	if len(q.Items) == 0 {
		return ""
	}
	item = q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return
}
