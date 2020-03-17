package stack

import (
	"sync"
)

type StackArr struct {
	Items []string
	mutx  *sync.RWMutex
}

//初始化操作
func (s *StackArr) Init() {
	s.Items = []string{}
	s.mutx = &sync.RWMutex{}
}

//入栈
func (s *StackArr) Push(item string) bool {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	s.Items = append(s.Items, item)
	return true
}

//出栈
func (s *StackArr) Pop() (item string) {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if len(s.Items) == 0 {
		return ""
	}
	item = s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return
}
