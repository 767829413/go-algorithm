package applyleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	assert := assert.New(t)
	head, _ := NewListNode([]int{1, 2, 3, 4, 5})
	expected := []int{5, 4, 3, 2, 1}
	assert.Equal(expected, reverseList(head).GetValueArray())
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
	head, _ := NewListNode([]int{1, 2, 3, 4, 5})
	k := 2
	expected := []int{2, 1, 4, 3, 5}
	assert.Equal(expected, reverseKGroup(head, k).GetValueArray())
}

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 0, 0}
	expected := [][]int{
		{0, 0, 0},
	}
	assert.Equal(expected, threeSum(nums))
}

func TestSortArray(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 2, 3, 1}
	expected := []int{1, 2, 3, 5}
	sortArray(nums)
	assert.Equal(expected, nums)
}

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	assert.Equal(expected, maxSubArray(input))
}

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)
	l1, _ := NewListNode([]int{1, 2, 4})
	l2, _ := NewListNode([]int{1, 3, 4})
	expected := []int{1, 1, 2, 3, 4, 4}
	assert.Equal(expected, mergeTwoLists(l1, l2).GetValueArray())
}

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 7, 11, 15}
	target := 18
	expected := []int{1, 2}
	assert.Equal(expected, twoSum(nums, target))
}

func TestLevelOrder(t *testing.T) {
	assert := assert.New(t)
	root := []int{1, 2, 3, 4, -1, -1, 5}
	expected := [][]int{
		{1},
		{2, 3},
		{4, 5},
	}
	node, _ := NewTreeNode(root)
	assert.Equal(expected, levelOrder(node))
}

func TestMaxProfit(t *testing.T) {
	assert := assert.New(t)
	input := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	assert.Equal(expected, maxProfit(input))
	assert.Equal(expected, maxProfit2(input))
}

func TestHasCycle(t *testing.T) {
	assert := assert.New(t)
	head, _ := NewListNode([]int{3, 2, 0, -4})
	head.Next.Next.Next.Next = head.Next
	expected := true
	assert.Equal(expected, hasCycle(head))
}

func TestNumIslands(t *testing.T) {
	assert := assert.New(t)
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	assert.Equal(expected, numIslands(grid))
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 1, 2, 3, 4}
	target := 1
	expected := 1
	assert.Equal(expected, search(nums, target))
}

func TestZigzagLevelOrder(t *testing.T) {
	assert := assert.New(t)
	root := []int{1, 2, 3, 4, -1, -1, 5}
	expected := [][]int{
		{1},
		{3, 2},
		{4, 5},
	}
	node, _ := NewTreeNode(root)
	assert.Equal(expected, zigzagLevelOrder(node))
}

func TestGetIntersectionNode(t *testing.T) {
	assert := assert.New(t)
	headA, mA := NewListNode([]int{1, 9, 1})
	headB, mB := NewListNode([]int{3, 5})
	var expected *ListNode
	Intersection, _ := NewListNode([]int{2, 4, 6})
	mA[2].Next = Intersection
	mB[1].Next = Intersection
	expected = Intersection
	assert.Equal(expected, getIntersectionNode(headA, headB))
	// assert.Equal(expected, getIntersectionNode(headA, headB))
}

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	s := "()"
	expected := true
	assert.Equal(expected, isValid(s))
}

func TestLowestCommonAncestor(t *testing.T) {
	assert := assert.New(t)
	root, m := NewTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	expected := m[1]
	assert.Equal(expected, lowestCommonAncestor(root, m[1], m[10]))
}

func TestLongestPalindrome(t *testing.T) {
	assert := assert.New(t)
	s := "babad"
	expected := "bab"
	assert.Equal(expected, longestPalindrome(s))
}

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	expected := []int{1, 2}
	merge(nums1, m, nums2, n)
	assert.Equal(expected, nums1)
}

func TestPermute(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	assert.Equal(expected, permute(nums))
}

func TestMergeKLists(t *testing.T) {
	assert := assert.New(t)
	list1, _ := NewListNode([]int{1, 4, 5})
	list2, _ := NewListNode([]int{1, 3, 4})
	list3, _ := NewListNode([]int{2, 6})
	lists := []*ListNode{list1, list2, list3}
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	assert.Equal(expected, mergeKLists(lists).GetValueArray())
}

func TestSpiralOrder(t *testing.T) {
	assert := assert.New(t)
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	assert.Equal(expected, spiralOrder(matrix))
}

func TestReverseBetween(t *testing.T) {
	assert := assert.New(t)
	head, _ := NewListNode([]int{3, 5})
	left, right := 1, 2
	expected := []int{5, 3}
	assert.Equal(expected, reverseBetween(head, left, right).GetValueArray())
}

func TestDetectCycle(t *testing.T) {
	assert := assert.New(t)
	head, m := NewListNode([]int{3, 2, 0, -4})
	pos := 1
	m[3].Next = m[pos]
	expected := m[pos]
	assert.Equal(expected, detectCycle(head))
}

func TestAddStrings(t *testing.T) {
	assert := assert.New(t)
	num1, num2 := "9", "1"
	expected := "10"
	assert.Equal(expected, addStrings(num1, num2))
}

func TestLengthOfLIS(t *testing.T) {
	assert := assert.New(t)
	nums := []int{}
	expected := 0
	assert.Equal(expected, lengthOfLIS(nums))
}

func TestTrap(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	assert.Equal(expected, trap(nums))
}
