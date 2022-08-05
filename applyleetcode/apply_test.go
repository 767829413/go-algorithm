package applyleetcode

import (
	"fmt"
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
	fmt.Println(cache.data, cache.head.Next, cache.tail.Pre)
	cache.Put(2, 6)
	fmt.Println(cache.data, cache.head.Next, cache.tail.Pre)
	assert.Equal(-1, cache.Get(1))
	fmt.Println(cache.data, cache.head.Next, cache.tail.Pre)
	cache.Put(1, 5)
	fmt.Println(cache.data, cache.head.Next, cache.tail.Pre)
	cache.Put(1, 2)
	fmt.Println(cache.data, cache.head.Next, cache.tail.Pre)
	assert.Equal(2, cache.Get(1))
	fmt.Println(cache.data, cache.head, cache.tail)
	assert.Equal(6, cache.Get(2))
	fmt.Println(cache.data, cache.head, cache.tail)
}
