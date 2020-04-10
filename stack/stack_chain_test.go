package stack

import (
	"fmt"
	"algorithm/linkedlist"
	"algorithm/utilcom/randomstring"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	//构建初始数据
	s := &StackChain{}
	s.Init(5)
	for i := 0; i < s.Count; i++ {
		node := &linkedlist.Node{}
		node.Item = randomstring.RandStringBytesMaskImprSrc(2)
		s.Push(node)
	}
	s.Items.Print()
	for s.Items.Size != 0 {
		fmt.Println("pop node")
		fmt.Println(s.Pop())
	}
	s.Items.Print()
}
