package linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		value: v,
	}
}

func (ln *ListNode) Next() *ListNode {
	return ln.next
}

func (ln *ListNode) Value() interface{} {
	return ln.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
	}
}

//在某个节点后面插入
func (ls *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	ls.length++
	return true
}

//在某个节点前面插入
func (ls *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	pre := ls.head
	for {
		if pre.next == p {
			break
		}
		pre = pre.next
		if pre == nil {
			return false
		}
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = p
	ls.length++
	return true
}

//在链表头部插入
func (ls *LinkedList) InsertToHead(v interface{}) bool {
	if ls.length == 0 {
		ls.head = NewListNode(v)
	} else {
		newNode := NewListNode(v)
		newNode.next = ls.head
		ls.head = newNode
	}
	ls.length++
	return true
}

//在链表尾部插入
func (ls *LinkedList) InsertToTail(v interface{}) bool {
	if ls.length == 0 {
		ls.head = NewListNode(v)
		ls.length++
		return true
	} else {
		cur := ls.head
		for cur.next != nil {
			cur = cur.next
		}
		return ls.InsertAfter(cur, v)
	}

}

//通过索引查找节点
func (ls *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= ls.length {
		return nil
	}
	cur := ls.head
	i := uint(0)
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (ls *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	pre := ls.head
	for {
		if pre.next == p {
			break
		}
		pre = pre.next
		if pre == nil {
			return false
		}
	}
	pre.next = p.next
	ls.length--
	return true
}

//打印链表
func (ls *LinkedList) Print() string {
	cur := ls.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Value())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	return format
}
