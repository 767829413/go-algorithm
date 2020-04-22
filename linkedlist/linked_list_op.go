package linkedlist

//反转链表
func ReverseLinkedList(l *LinkedList) bool {
	if l.length == 1 || l.length == 0 {
		return false
	}
	pre := l.head
	cur := l.head.next
	l.head.next = nil
	var tail *ListNode
	for cur != nil {
		temp := cur.next
		if temp == nil {
			tail = cur
		}
		cur.next = pre
		pre = cur
		cur = temp
	}
	l.head = tail
	return true
}

//判断单链表是否有环
func HasLoop(l *LinkedList) bool {
	if l.length == 1 || l.length == 0 {
		return false
	}
	low := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		low = low.next
		fast = fast.next.next
		if low == fast {
			return true
		}
	}
	return false
}

//合并有序链表
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil {
		return l2
	}
	if l2 == nil || l2.head == nil {
		return l1
	}
	l := NewLinkedList()
	cur1 := l1.head
	cur2 := l2.head
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			l.InsertToTail(cur2.value)
			cur2 = cur2.next
		} else {
			l.InsertToTail(cur1.value)
			cur1 = cur1.next
		}
	}
	if cur1 != nil {
		l.InsertToTail(cur1.value)
	} else if cur2 != nil {
		l.InsertToTail(cur2.value)
	}
	return l
}

//删除倒数第N个节点
func DeleteBottomN(l *LinkedList, index uint) bool {
	if index > l.length {
		return false
	}
	if index == l.length {
		l.head = nil
		l.length--
		return true
	}
	needIndex := l.length - index
	pre := l.head
	cur := l.head.next
	i := uint(0)
	for {
		i++
		if i == needIndex {
			break
		}
		cur = cur.next
		pre = pre.next
	}
	if cur.next == nil {
		pre.next = nil
	} else {
		pre.next = cur.next
	}
	l.length--
	return true
}

//找到中点
func FindMiddleNode(l *LinkedList) *ListNode {
	if l.length == 0 {
		return nil
	}
	if l.length == 1 {
		return l.head
	}
	low := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		low = low.next
	}
	return low
}
