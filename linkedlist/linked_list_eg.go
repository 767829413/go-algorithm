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

func (this *ListNode) Next() *ListNode {
	return this.next
}

func (this *ListNode) Value() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
	}
}

//在某个节点后面插入
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个节点前面插入
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	pre := this.head
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
	this.length++
	return true
}

//在链表头部插入
func (this *LinkedList) InsertToHead(v interface{}) bool {
	if this.length == 0 {
		this.head = NewListNode(v)
	} else {
		newNode := NewListNode(v)
		newNode.next = this.head
		this.head = newNode
	}
	this.length++
	return true
}

//在链表尾部插入
func (this *LinkedList) InsertToTail(v interface{}) bool {
	if this.length == 0 {
		this.head = NewListNode(v)
		this.length++
		return true
	} else {
		cur := this.head
		for cur.next != nil {
			cur = cur.next
		}
		return this.InsertAfter(cur, v)
	}

}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head
	i := uint(0)
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	pre := this.head
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
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() string {
	cur := this.head
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
