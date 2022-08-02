package applyleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Values []int) *ListNode {
	l := len(Values)
	if l == 0 {
		panic("Wrong input")
	}
	var recurrence func(start int) *ListNode
	recurrence = func(start int) *ListNode {
		if start == l {
			return nil
		}
		node := &ListNode{Val: Values[start]}
		node.Next = recurrence(start + 1)
		return node
	}
	return recurrence(0)
}

func (ln *ListNode) GetValueArray() []int {
	res := []int{}
	cur := ln
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
