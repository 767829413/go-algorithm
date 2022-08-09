package applyleetcode

import (
	"math"
	"sort"
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
	heapSize := l
	// 构建大顶堆
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 这里是交换最大最小值,来进行一次自顶向下的全堆化
		nums[0], nums[i] = nums[i], nums[0]
		// 这是缩减堆的大小,逼近所期望的第k大值
		heapSize--
		// 进行对话,堆顶是最大值
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// Reverse nodes in k group
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 设置哨兵,方便操作
	sentinel := &ListNode{Val: 0, Next: head}
	// 构建快慢指针节点,进行k对翻转
	pre, end := sentinel, sentinel
	// 链表翻转操作
	reverse := func(start *ListNode) *ListNode {
		var pre, cur *ListNode
		pre, cur = nil, start
		for cur != nil {
			tmp := cur.Next
			cur.Next = pre
			pre, cur = cur, tmp
		}
		return pre
	}
	for end.Next != nil {
		// 快指针节点进行k次向前
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果快指针节点为nil,那么就是看无法被链表整除,剩下的不用翻转
		if end == nil {
			break
		}
		// 构建临时变量记录快,慢指针节点的下一位
		start, next := pre.Next, end.Next
		// 这里讲快指针节点下一位置空是为了方便翻转快慢指针之间的节点
		end.Next = nil
		// 执行翻转操作
		pre.Next = reverse(start)
		// 翻转后更新位置
		start.Next, pre = next, start
		end = pre
	}

	return sentinel.Next
}

// 3sum
func threeSum(nums []int) [][]int {
	l, res := len(nums), [][]int{}
	if l <= 3 {
		if l == 3 && nums[0]+nums[1]+nums[2] == 0 {
			return append(res, nums)
		}
		return res
	}

	// 排序,加速判断
	sort.Ints(nums)
	// 遍历元素
	for i := range nums {
		// 利用排序的结果快速过滤
		if nums[i] > 0 {
			break
		}
		// 过滤重复
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		//准备两个指针来组合结果
		lo, hi := i+1, l-1
		for lo < hi {
			r := nums[i] + nums[lo] + nums[hi]
			if r == 0 {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				// 过滤重复
				for lo < hi && nums[lo+1] == nums[lo] {
					lo++
				}
				// 过滤重复
				for lo < hi && nums[hi-1] == nums[hi] {
					hi--
				}
				lo++
				hi--
			} else if r > 0 {
				// 过滤重复
				for lo < hi && nums[hi-1] == nums[hi] {
					hi--
				}
				hi--
			} else {
				// 过滤重复
				for lo < hi && nums[lo+1] == nums[lo] {
					lo++
				}
				lo++
			}
		}
	}
	return res
}
