package queue

import (
	"fmt"
	"sync"
)

/**
基于数组的循环队列
会浪费一个放数据的位置
 */

type CircularQueue struct {
	data []interface{}
	cap  int //容量
	head int //头指针
	tail int //尾指针
	mut  *sync.Mutex
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{
		data: make([]interface{}, n),
		cap:  n,
		head: 0,
		tail: 0,
		mut:  &sync.Mutex{},
	}
}

//循环队列为空的条件 tail == head
func (cq *CircularQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

//循环队列 满 的条件 (tail+1) % capacity == head
func (cq *CircularQueue) IsFull() bool {
	return ((cq.tail + 1) % cq.cap) == cq.head
}

//入队
func (cq *CircularQueue) EnQueue(v interface{}) bool {
	cq.mut.Lock()
	defer cq.mut.Unlock()
	if cq.IsFull() {
		return false
	}
	cq.data[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.cap
	return true
}

//出队
func (cq *CircularQueue) DeQueue() interface{} {
	cq.mut.Lock()
	defer cq.mut.Unlock()
	if cq.IsEmpty() {
		return nil
	}
	v := cq.data[cq.head]
	cq.head = (cq.head + 1) % cq.cap
	return v
}

func (cq *CircularQueue) Print() (format string) {
	cq.mut.Lock()
	defer cq.mut.Unlock()
	if cq.IsEmpty() {
		format = "empty"
	} else {
		format = "head"
		for i := cq.head; i != cq.tail; i = (i + 1) % cq.cap {
			format += fmt.Sprintf("<-%+v", cq.data[i])
		}
		format += "<-tail"
	}
	return
}
