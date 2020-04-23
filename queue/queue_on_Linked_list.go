package queue

import "fmt"

/**
基于链表的非循环队列
 */
type listNode struct {
	v    interface{}
	next *listNode
}

type LinkedListQueue struct {
	head   *listNode
	tail   *listNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (llq *LinkedListQueue) EnQueue(v interface{}) {
	node := &listNode{
		v:    v,
		next: nil,
	}
	if llq.tail == nil {
		llq.tail = node
		llq.head = node
	} else {
		llq.tail.next = node
		llq.tail = node
	}
	llq.length++
}

func (llq *LinkedListQueue) DeQueue() interface{} {
	if llq.head == nil {
		return nil
	}
	v := llq.head.v
	llq.head = llq.head.next
	llq.length--
	return v
}

func (llq *LinkedListQueue) Print() (format string) {
	if llq.head == nil {
		format = "empty"
	} else {
		format = "head"
		for cur := llq.head; cur != nil; cur = cur.next {
			format += fmt.Sprintf("<-%+v", cur.v)
		}
		format += "<-tail"
	}
	return
}
