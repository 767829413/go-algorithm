package queue

import (
	"fmt"
	"testing"

	"github.com/767829413/go-algorithm/linkedlist"
	"github.com/767829413/go-algorithm/utilcom"
)

func TestQueueChainEnDeQueue(t *testing.T) {
	//构建测试顺序队列
	q := NewQueueChain(6)
	for i := 0; i < q.Count; i++ {
		node := &linkedlist.Node{}
		node.Item = utilcom.RandStringBytesMaskImprSrc(3)
		q.EnQueue(node)
	}
	q.Items.Print()
	fmt.Println("----")
	for q.Items.Size != 0 {
		fmt.Println(q.DeQueue())
	}
	q.Items.Print()
}
