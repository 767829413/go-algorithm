package stack

import (
	"sync"
)

type Stack struct {
	Items []string
	Count int
	mutx  *sync.RWMutex
}

//初始化操作
func (s *Stack) Init(n int) {
	s.Count = 0
	s.Items = []string{}
	s.mutx = &sync.RWMutex{}
}

//入栈
func (s *Stack) Push(item string) bool {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	s.Items = append(s.Items, item)
	s.Count++
	return true
}

//出栈
func (s *Stack) Pop() (item string) {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if s.Count == 0 {
		return ""
	}
	item = s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Count--
	return
}
