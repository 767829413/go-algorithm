package stack

import (
	"fmt"
	"go-algorithm/linkedlist"
	"go-algorithm/utilcom/randomstring"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	//构建初始数据
	s := &StackChain{}
	s.Init()
	for i := 0; i < 10; i++ {
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
