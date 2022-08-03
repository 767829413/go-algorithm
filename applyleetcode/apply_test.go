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
