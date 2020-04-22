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

func (this *LinkedListStack) IsEmpty() bool {
	return this.TopNode == nil
}

func (this *LinkedListStack) Push(v interface{}) {
	this.TopNode = &node{next: this.TopNode, val: v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.TopNode
	this.TopNode = this.TopNode.next
	return v
}

func (this *LinkedListStack) Top() interface{} {
	return this.TopNode.val
}

func (this *LinkedListStack) Flush() {
	this.TopNode = nil
}

func (this *LinkedListStack) Print() (format string) {
	if this.IsEmpty() {
		format = "empty"
	} else {
		cur := this.TopNode
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
