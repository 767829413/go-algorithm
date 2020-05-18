package queue

import "fmt"

/**
基于数组的非循环队列
 */
type ArrayQueue struct {
	data []interface{}
	cap  int //容量
	head int //头指针
	tail int //尾指针
}

func NewArrayQueue(n int) *ArrayQueue {
	if n == 0 {
		return nil
	}
	return &ArrayQueue{
		data: make([]interface{}, n),
		cap:  n,
		head: 0,
		tail: 0,
	}
}

func (aq *ArrayQueue) Full() bool {
	return aq.tail == aq.cap
}

func (aq *ArrayQueue) Empty() bool {
	return aq.head == aq.tail
}

func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.cap {
		return false
	}
	aq.data[aq.tail] = v
	aq.tail++
	return true
}

func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.data[aq.head]
	aq.head++
	return v
}

func (aq *ArrayQueue) Print() (format string) {
	if aq.head == aq.tail {
		format = "empty"
	} else {
		format = "head"
		for i := aq.head; i < aq.tail; i++ {
			format += fmt.Sprintf("<-%+v", aq.data[i])
		}
		format += "<-tail"
	}
	return
}
