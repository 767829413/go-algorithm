package stack

import (
	"fmt"
	"go-algorithm/utilcom"
	"testing"
)

func TestPushPop(t *testing.T) {
	s := &Stack{}
	s.Init(10)
	for i := 0; i < 100; i++ {
		if !s.Push(utilcom.RandStringBytesMaskImprSrc(3)) {
			t.Error("push err")
		}
	}
	fmt.Println(s)
	for ; s.Count != 0; {
		fmt.Println(s.Pop())
	}
}
