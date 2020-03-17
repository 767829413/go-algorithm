package stack

import (
	"go-algorithm/linkedlist"
)

type StackChain struct {
	Items linkedlist.SingleList
}

//初始化操作
func (s *StackChain) Init() {
	s.Items = linkedlist.SingleList{}
	s.Items.Init()
}

//入栈
func (s *StackChain) Push(node *linkedlist.Node) bool {
	s.Items.Append(node)
	return true
}

//出栈
func (s *StackChain) Pop() (node *linkedlist.Node) {
	if s.Items.Size == 0 {
		node = nil
		return
	}
	node = s.Items.Head
	s.Items.Delete(0)
	return
}
