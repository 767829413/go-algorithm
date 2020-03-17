package queue

import (
	"sync"
)

type QueueArr struct {
	Mutx  *sync.RWMutex
	Count int
	Items []string
}

func NewQueueArr(n int) *QueueArr {
	return &QueueArr{
		Mutx:  &sync.RWMutex{},
		Count: n,
		Items: []string{},
	}
}

//入队
func (q *QueueArr) EnQueue(item string) bool {
	if q.Count == len(q.Items) {
		return false
	}
	q.Items = append(q.Items, item)
	return true
}

//出队
func (q *QueueArr) DeQueue() (item string) {
	if len(q.Items) == 0 {
		return ""
	}
	item = q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return
}
