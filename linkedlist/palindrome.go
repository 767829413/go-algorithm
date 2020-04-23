package linkedlist

/**
判断是否是回文
思路1  : 通过一个栈来存放链表前半段
思路2 : 通过翻转前半部分对比来判断
 */

func IsPalindromeByStack(l *LinkedList) bool {
	if l.length == 0 {
		return false
	}
	if l.length == 1 {
		return true
	}
	mid := l.length >> 1
	isResidue := l.length%2 == 0
	//利用切片构建栈
	s := make([]interface{}, 0, mid)
	cur := l.head
	i := uint(0)
	for ; i < l.length; i++ {
		if !isResidue && i == mid {
			cur = cur.next
			continue
		}
		if i < mid {
			s = append(s, cur.value)
		} else {
			if s[l.length-i-1] != cur.value {
				return false
			}
		}
		cur = cur.next
	}
	return true
}

func IsPalindromeByFastLowPoint(l *LinkedList) bool {
	if l.length == 0 {
		return false
	}
	if l.length == 1 {
		return true
	}
	if l.length == 2 {
		if l.head.value != l.head.next.value {
			return false
		}
		return true
	}
	var mid uint
	if l.length <= 5 {
		mid = l.length >> 2
	} else {
		mid = (l.length >> 2) + 1
	}
	pre := l.head
	cur := l.head.next
	l.head.next = nil
	i := uint(0)
	//折半翻转
	for i = 0; i < mid; i++ {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	var toLeft *ListNode
	var toRight *ListNode
	if l.length%2 != 0 {
		toLeft = pre
		toRight = cur.next
	} else {
		toLeft = pre
		toRight = cur
	}
	var isPalindrome = true
	for toLeft != nil && toRight != nil {
		if toLeft.value != toRight.value {
			isPalindrome = false
			break
		}
		toLeft = toLeft.next
		toRight = toRight.next
	}
	//还原链表
	pre, cur = cur, pre
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return isPalindrome
}
