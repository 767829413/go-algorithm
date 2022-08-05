package applyleetcode

import (
	"math"
)

// Reverse linked list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre, cur = nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}

// Longest substring without repeating characters
// 这类匹配子类问题的通解一般是滑动窗口
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}
	res, a, m, max := math.MinInt, []byte(s), make(map[byte]int, l), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, j := 0, 0; j < l; j++ {
		// fmt.Println(res, i, j, m)
		if v, ok := m[a[j]]; ok {
			// v这里的是重复元素最初出现的位置
			// 和滑动窗口左侧比较(可能位置0后续重复)
			i = max(i, v)
		}
		// 更新或记录每个值的不重复位置
		m[a[j]] = j + 1
		res = max(res, j-i+1)
	}
	return res
}

// Kth largest element in an array
func findKthLargest(nums []int, k int) int {
	l := len(nums)
	if k > l || k < 0 {
		return -1
	}
	// sort.Ints(nums)
	// MergeSort(nums)
	// QuickSort(nums)
	// return nums[l-k]
	// 堆排序
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}
