package stack

import (
	"go-algorithm/linkedlist"
)

type StackChain struct {
	Items linkedlist.SingleList
	Count int
}

//初始化操作
func (s *StackChain) Init(n int) {
	s.Items = linkedlist.SingleList{}
	s.Count = n
	s.Items.Init()
}

//入栈
func (s *StackChain) Push(node *linkedlist.Node) bool {
	if s.Count == s.Items.Size {
		return false
	}
	s.Items.Append(node)
	return true
}

//出栈
func (s *StackChain) Pop() (node *linkedlist.Node) {
	if s.Items.Size == 0 {
		node = nil
		return
	}
	node = s.Items.Tail
	s.Items.Delete(s.Items.Size - 1)
	return
}
