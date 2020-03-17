package queue

import (
	"fmt"
	"go-algorithm/utilcom/randomstring"
	"testing"
)

func TestEnAndDeQueue(t *testing.T) {
	//构建测试顺序队列
	q := NewQueue(100)
	for i := 0; i < q.Count; i++ {
		q.EnQueue(randomstring.RandStringBytesMaskImprSrc(3))
		fmt.Println(q)
	}
	fmt.Println("----")
	for ; len(q.Items) != 0; {
		fmt.Println(q.DeQueue())
	}
	fmt.Println(q)
}
