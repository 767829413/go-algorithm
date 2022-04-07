package linkedlist

import (
	"testing"

	"github.com/767829413/go-algorithm/utilcom/randomstring"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(randomstring.RandStringBytesMaskImprSrc(1))
	}
	t.Log(l.Print())
	ReverseLinkedList(l)
	t.Log(l.Print())
}

func TestIsHasLoop(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 4; i++ {
		l.InsertToTail(randomstring.RandStringBytesMaskImprSrc(1))
	}
	t.Log(HasLoop(l))
	l.head.next.next.next = l.head
	t.Log(HasLoop(l))
}

func TestMergeSortedList(t *testing.T) {
	l1 := NewLinkedList()
	l2 := NewLinkedList()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			l1.InsertToTail(i)
		} else {
			l2.InsertToTail(i)
		}
	}
	t.Log(l1.Print())
	t.Log(l2.Print())
	t.Log(MergeSortedList(l1, l2).Print())
}

func TestDeleteBottomN(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 4; i++ {
		l.InsertToTail(randomstring.RandStringBytesMaskImprSrc(1))
	}
	t.Log(l.Print())
	assert.Equal(t, false, DeleteBottomN(l, 100), l.Print())
	t.Log(l.Print())
	assert.Equal(t, true, DeleteBottomN(l, 1), l.Print())
	t.Log(l.Print())
	assert.Equal(t, true, DeleteBottomN(l, 1), l.Print())
	t.Log(l.Print())
	assert.Equal(t, true, DeleteBottomN(l, 1), l.Print())
	t.Log(l.Print())
	assert.Equal(t, true, DeleteBottomN(l, 1), l.Print())
	t.Log(l.Print())
	assert.Equal(t, false, DeleteBottomN(l, 1), l.Print())
	t.Log(l.Print())
}

func TestFindMiddleNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(randomstring.RandStringBytesMaskImprSrc(1))
	}
	t.Log(l.Print())
	t.Log(FindMiddleNode(l).value)
}
