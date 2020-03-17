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
	pre := sl.Head
	for i := 0; i < index-1; i++ {
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
	if sl.Size == 0 {
		fmt.Println(nil)
	}
	if sl.Head != nil {
		pre := sl.Head
		for i := 0; i < sl.Size && pre != nil; i++ {
			fmt.Println(pre)
			pre = pre.Next
		}
	}
}

//是否有环
func (sl *SingleList) HasLoop() bool {
	if sl.Head != nil && sl.Head.Next != nil {
		low := sl.Head
		fast := sl.Head.Next
		for {
			low = low.Next
			if fast.Next == nil {
				break
			}
			fast = fast.Next.Next
			if fast == nil {
				break
			}
			if low.Item.(int) == fast.Item.(int) {
				return true
			}
		}
	}
	return false
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
			for curNode != nil {
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

//单链表合并
func (sl *SingleList) Merge(s *SingleList, option int) {
	if sl.Head == nil {
		sl.Init()
		sl = s
		return
	}
	if s.Head == nil {
		return
	}
	if sl.Tail.Item.(int) < s.Tail.Item.(int) {
		sl.Tail = s.Tail
	}
	sl.Size = sl.Size + s.Size
	switch option {
	case IsRecursive:
		sl.Head = sl.MergeRecursive(sl.Head, s.Head)
	default:
		sl.MergeNotRecursive(sl.Head, s.Head)
	}
	s.Init()
}

//有序单链表合并 - 递归
func (sl *SingleList) MergeRecursive(headBe *Node, headAf *Node) (newHead *Node) {
	if headBe == nil {
		return headAf
	}
	if headAf == nil {
		return headBe
	}
	if headAf.Item.(int) > headBe.Item.(int) {
		newHead = headBe
		newHead.Next = sl.MergeRecursive(headBe.Next, headAf)
	} else {
		newHead = headAf
		newHead.Next = sl.MergeRecursive(headBe, headAf.Next)
	}
	return newHead
}

//有序单链表合并 - 非递归
func (sl *SingleList) MergeNotRecursive(headBe *Node, headAf *Node) {
	var head *Node
	if headBe == nil {
		return
	}
	if headAf == nil {
		return
	}
	if headBe.Item.(int) > headAf.Item.(int) {
		head = headAf
		headAf = headAf.Next
	} else {
		head = headBe
		headBe = headBe.Next
	}
	sl.Head = head
	for {
		if headAf == nil {
			if head.Next == nil {
				head.Next = headBe
			} else {
				head.Next.Next = headBe
			}
			break
		}
		if headBe == nil {
			if head.Next == nil {
				head.Next = headAf
			} else {
				head.Next.Next = headAf
			}
			break
		}
		if headBe.Item.(int) > headAf.Item.(int) {
			head.Next = headAf
			headAf = headAf.Next
		} else {
			head.Next = headBe
			headBe = headBe.Next
		}
		head = head.Next
	}
}

//删除链表倒数第 index 节点
func (sl *SingleList) DelByReciprocal(index int) bool {
	if sl.Size-index < 0 {
		return false
	}
	if sl.Size-index == 0 {
		sl.Head = sl.Head.Next
		sl.Size--
		return true
	}
	fast := sl.Head
	low := sl.Head
	for i := 0; i < index; i++ {
		fast = fast.Next
	}
	if fast == nil {
		low.Next = fast
		sl.Size--
		return true
	}
	for ; fast.Next != nil; {
		fast = fast.Next
		low = low.Next
	}
	low.Next = low.Next.Next
	sl.Size--
	return true
}

//求链表中间节点
func (sl *SingleList) FindMiddle() *Node {
	fast := sl.Head
	low := sl.Head
	for ; fast != nil; {
		if fast.Next == nil {
			break
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}
