package tree

import (
	"algorithm/utilcom/randomstring"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()
	num := 12
	testData := make(map[int]string, num)
	for i := 0; i < num; i++ {
		index := rand.Intn(20)
		v := randomstring.RandStringBytesMaskImprSrc(2)
		testData[index] = v
		tree.Insert(v, index)

	}
	tree.Insert("max", 21)
	tree.Insert("min", 0)
	testData[0] = "min"
	testData[21] = "max"
	assert.Equal(t, "max", tree.FindMax())
	assert.Equal(t, "min", tree.FindMin())
	for k, v := range testData {
		assert.Equal(t, v, tree.Find(k), k, v)
		tree.Delete(k)
		assert.Equal(t, nil, tree.Find(k), k, v)
	}
}
