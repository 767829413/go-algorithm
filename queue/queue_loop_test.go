package queue

import (
	"fmt"
	"testing"

	"github.com/767829413/go-algorithm/utilcom/randomstring"
)

func TestQueueLoopEnDe(t *testing.T) {
	//构建测试顺序队列
	q := NewQueueLoop(5)
	for i := 0; i < 5; i++ {
		q.EnQueue(randomstring.RandStringBytesMaskImprSrc(3))
		fmt.Println(q)
	}
	fmt.Println("----")
	for {
		str := q.DeQueue()
		if str == "" {
			break
		}
	}
	fmt.Println("----")
	for i := 0; i < 5; i++ {
		q.EnQueue(randomstring.RandStringBytesMaskImprSrc(3))
		fmt.Println(q)
	}
}
