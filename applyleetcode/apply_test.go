package applyleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	assert := assert.New(t)
	head := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	assert.Equal(expected, reverseList(NewListNode(head)).GetValueArray())
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)
	// s := "pwwkew"
	s := "dvdf"
	expected := 3
	assert.Equal(expected, lengthOfLongestSubstring(s))
}

func TestLRUCache(t *testing.T) {
	assert := assert.New(t)
	cache := Constructor(2)
	assert.Equal(-1, cache.Get(2))
	cache.Put(2, 6)
	assert.Equal(-1, cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	assert.Equal(2, cache.Get(1))
	assert.Equal(6, cache.Get(2))
}

func TestFindKthLargest(t *testing.T) {
	assert := assert.New(t)
	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	expected := 4
	assert.Equal(expected, findKthLargest(input, k))
}

func TestReverseKGroup(t *testing.T) {
	assert := assert.New(t)
	head := []int{1, 2, 3, 4, 5}
	k := 2
	expected := []int{2, 1, 4, 3, 5}
	assert.Equal(expected, reverseKGroup(NewListNode(head), k).GetValueArray())
}

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 0, 0}
	expected := [][]int{
		{0, 0, 0},
	}
	assert.Equal(expected, threeSum(nums))
}
