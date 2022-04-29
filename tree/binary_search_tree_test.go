package tree

import (
	"math/rand"
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
	"github.com/stretchr/testify/assert"
)

func TestNewBinarySearchTree_Inset_Find(t *testing.T) {
	tatil := 1000
	for i := 0; i < tatil; i++ {
		tree := NewBinarySearchTree()
		num := 12
		testData := make(map[int]string, num)
		for i := 0; i < num; i++ {
			index := rand.Intn(20)
			v := utilcom.RandStringBytesMaskImprSrc(2)
			testData[index] = v
			tree.Insert(v, index)
		}
		for k, v := range testData {
			assert.Equal(t, v, tree.Find(k))
		}
	}
}

func TestNewBinarySearchTree_Find_Max_Min(t *testing.T) {
	num := 120
	testData := make(map[int]string, num)
	tree := NewBinarySearchTree()
	min := 21
	minV := ""
	max := -1
	maxV := ""
	for i := 0; i < num; i++ {
		index := rand.Intn(20)
		if _, ok := testData[index]; ok {
			continue
		}
		v := utilcom.RandStringBytesMaskImprSrc(2)
		if min > index {
			min = index
			minV = v
		}
		if max < index {
			max = index
			maxV = v
		}
		testData[index] = v
		tree.Insert(v, index)
	}
	assert.Equal(t, maxV, tree.FindMax())
	assert.Equal(t, minV, tree.FindMin())
}

func TestNewBinarySearchTree_Delete(t *testing.T) {
	tatil := 10000
	for i := 0; i < tatil; i++ {
		tree := NewBinarySearchTree()
		num := 20
		testData := make(map[int]string, num)
		for i := 0; i < num; i++ {
			index := rand.Intn(20)
			v := utilcom.RandStringBytesMaskImprSrc(2)
			testData[index] = v
		}
		for k, v := range testData {
			tree.Insert(v, k)
		}
		delKeys := make([]int, 0)
		curKeys := make([]int, 0)
		for k, _ := range testData {
			if k < 10 {
				curKeys = append(curKeys, k)
			} else {
				delKeys = append(delKeys, k)
				tree.Delete(k)
			}
		}
		for _, index := range delKeys {
			assert.Equal(t, nil, tree.Find(index))
		}
		for _, index := range curKeys {
			assert.Equal(t, testData[index], tree.Find(index))
		}
	}
}
