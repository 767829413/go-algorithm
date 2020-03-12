package linkedlist

import (
	"fmt"
	"sync"
)

const (
	IsRecursive = 0
)

type Node struct {
	Item interface{}
	Next *Node
}

type SingleList struct {
	Head  *Node
	Tail  *Node
	Size  int
	mutex *sync.RWMutex
}

func (sl *SingleList) Init() {
	sl.Head = nil
	sl.Tail = nil
	sl.Size = 0
	sl.mutex = &sync.RWMutex{}
}

//添加节点到尾部
func (sl *SingleList) Append(node *Node) bool {
	if node == nil {
		return false
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	if sl.Size == 0 {
		sl.Head = node
		sl.Tail = node
		sl.Size = 1
		return true
	}
	tail := sl.Tail
	tail.Next = node
	sl.Tail = node
	sl.Size += 1
	return true
}

//插入节点到指定位置
func (sl *SingleList) Insert(index int, node *Node) bool {
	if node == nil {
		return false
	}

	if index > sl.Size {
		return false
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	if index == 0 {
		node.Next = sl.Head
		sl.Head = node
		sl.Size += 1
		return true
	}
	pre := sl.Head
	for i := 1; i < index; i++ {
		pre = pre.Next
	}
	next := pre.Next
	pre.Next = node
	node.Next = next
	sl.Size += 1
	return false
}

//删除节点
func (sl *SingleList) Delete(index int) bool {
	tailIndex := sl.Size - 1
	if sl == nil || sl.Size == 0 || index > tailIndex {
		return false
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	if index == 0 {
		next := sl.Head.Next
		sl.Head = next
		if sl.Size == 1 {
			sl.Tail = nil
		}
		sl.Size -= 1
		return true
	}
	pre := sl.Head.Next
	for i := 1; i < index; i++ {
		pre = pre.Next
	}
	next := pre.Next
	pre.Next = next.Next
	if index == tailIndex {
		sl.Tail = pre
	}
	sl.Size -= 1
	return true
}

//查询指定节点
func (sl *SingleList) Find(index int) (node *Node) {
	tailIndex := sl.Size - 1
	if sl == nil || sl.Size == 0 || index > tailIndex {
		return nil
	}
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	if index == 0 {
		return sl.Head
	}
	if index == tailIndex {
		return sl.Tail
	}
	pre := sl.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	return pre
}

//打印链表
func (sl *SingleList) Print() {
	if sl.Head != nil {
		pre := sl.Head
		for i := 0; i < sl.Size; i++ {
			fmt.Println(pre)
			pre = pre.Next
		}
	}
	fmt.Println(nil)
}

//反转
func (sl *SingleList) Reverse(option int) {
	switch option {
	case IsRecursive:
		sl.ReverseByRecursive(sl.Head)
		head := sl.Head
		sl.Head = sl.Tail
		sl.Tail = head
	default:
		sl.ReverseByNotRecursive()
	}

}

//递归反转
func (sl *SingleList) ReverseByRecursive(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	successor := sl.ReverseByRecursive(node.Next)
	successor.Next = node
	node.Next = nil
	return node
}

//非递归反转
func (sl *SingleList) ReverseByNotRecursive() {
	if sl.Head != nil {
		if sl.Head.Next != nil {
			preNode := sl.Head
			curNode := sl.Head.Next
			tempNode := sl.Head.Next.Next
			for ; curNode != nil; {
				tempNode = curNode.Next
				curNode.Next = preNode
				preNode = curNode
				curNode = tempNode
			}
		}
		tail := sl.Tail
		sl.Head.Next = nil
		sl.Tail = sl.Head
		sl.Head = tail
	}
}
