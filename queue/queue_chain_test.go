package queue

import (
	"fmt"
	"go-algorithm/linkedlist"
	"go-algorithm/utilcom/randomstring"
	"testing"
)

func TestQueueChainEnDeQueue(t *testing.T) {
	//构建测试顺序队列
	q := NewQueueChain(6)
	for i := 0; i < q.Count; i++ {
		node := &linkedlist.Node{}
		node.Item = randomstring.RandStringBytesMaskImprSrc(3)
		q.EnQueue(node)
	}
	q.Items.Print()
	fmt.Println("----")
	for ; q.Items.Size != 0; {
		fmt.Println(q.DeQueue())
	}
	q.Items.Print()
}
