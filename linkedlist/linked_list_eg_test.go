package linkedlist

import (
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	t.Log(l.Print())
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.Print())
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(utilcom.RandStringBytesMaskImprSrc(1))
	}
	t.Log(l.Print())
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(utilcom.RandStringBytesMaskImprSrc(1))
	}
	t.Log(l.Print())

	t.Log(l.DeleteNode(l.head.next))
	t.Log(l.Print())

	t.Log(l.DeleteNode(l.head.next.next))
	t.Log(l.Print())
}
