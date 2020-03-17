package stack

import (
	"fmt"
	"go-algorithm/utilcom/randomstring"
	"testing"
)

func TestPushPop(t *testing.T) {
	s := &StackArr{}
	s.Init(10)
	for i := 0; i < s.Count; i++ {
		if !s.Push(randomstring.RandStringBytesMaskImprSrc(3)) {
			t.Error("push err")
		}
		fmt.Println(s)
	}
	for len(s.Items) != 0 {
		fmt.Println(s.Pop())
	}

}
