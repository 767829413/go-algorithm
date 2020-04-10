package queue

import (
	"algorithm/linkedlist"
)

type QueueChain struct {
	Items linkedlist.SingleList
	Count int
}

func NewQueueChain(n int) *QueueChain {
	queue := &QueueChain{
		Items: linkedlist.SingleList{},
		Count: n,
	}
	queue.Items.Init()
	return queue
}

//入队
func (q *QueueChain) EnQueue(node *linkedlist.Node) bool {
	if q.Count == q.Items.Size {
		return false
	}
	q.Items.Append(node)
	return true
}

//出队
func (q *QueueChain) DeQueue() (node *linkedlist.Node) {
	if q.Items.Size == 0 {
		return nil
	}
	node = q.Items.Head
	q.Items.Delete(0)
	return
}
