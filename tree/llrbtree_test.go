package tree

import (
	"algorithm/utilcom/randomstring"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLLRBTree(t *testing.T) {
	h := NewLLRBTree()
	num := 12
	testData := make(map[int]string, num)
	for i := 1; i < num; i++ {
		v := randomstring.RandStringBytesMaskImprSrc(2)
		testData[i] = v
		h.Inset(i, v)
	}
	for k, v := range testData {
		assert.Equal(t, v, h.Find(k).value)
	}
	for k, _ := range testData {
		h.Delete(k)
		assert.Equal(t, true, h.Find(k) == nil)
	}
}
