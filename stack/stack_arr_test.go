package stack

import (
	"fmt"
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
)

func TestPushPop(t *testing.T) {
	s := &StackArr{}
	s.Init(10)
	for i := 0; i < s.Count; i++ {
		if !s.Push(utilcom.RandStringBytesMaskImprSrc(3)) {
			t.Error("push err")
		}
		fmt.Println(s)
	}
	for len(s.Items) != 0 {
		fmt.Println(s.Pop())
	}

}
