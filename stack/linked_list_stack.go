package stack

import "fmt"

/**
基于链表实现
 */
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	TopNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{TopNode: nil}
}

func (lls *LinkedListStack) IsEmpty() bool {
	return lls.TopNode == nil
}

func (lls *LinkedListStack) Push(v interface{}) {
	lls.TopNode = &node{next: lls.TopNode, val: v}
}

func (lls *LinkedListStack) Pop() interface{} {
	if lls.IsEmpty() {
		return nil
	}
	v := lls.TopNode
	lls.TopNode = lls.TopNode.next
	return v
}

func (lls *LinkedListStack) Top() interface{} {
	return lls.TopNode.val
}

func (lls *LinkedListStack) Flush() {
	lls.TopNode = nil
}

func (lls *LinkedListStack) Print() (format string) {
	if lls.IsEmpty() {
		format = "empty"
	} else {
		cur := lls.TopNode
		format = ""
		for cur != nil {
			format += fmt.Sprintf("%+v", cur.val)
			format += "->"
			cur = cur.next
		}
		format += "end"
	}
	return
}
