package applyleetcode

import (
	"math"
	"sort"
	"strconv"
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

// Sort an array
func sortArray(nums []int) []int {
	// 已排序检查
	if isSorted(nums) {
		return nums
	}

	if false {
		// 快排
		position := func(nums []int, i, j int) int {
			pre, posVal := i, nums[j]
			for l := i; l < j; l++ {
				if nums[l] < posVal {
					nums[pre], nums[l] = nums[l], nums[pre]
					pre++
				}
			}
			nums[pre], nums[j] = nums[j], nums[pre]
			return pre
		}
		var quickSortPart func(nums []int, i, j int)
		quickSortPart = func(nums []int, i, j int) {
			if i >= j {
				return
			}
			pos := position(nums, i, j)
			quickSortPart(nums, i, pos-1)
			quickSortPart(nums, pos+1, j)
		}
		quickSort := func(nums []int) {
			l := len(nums)
			end := l - 1
			quickSortPart(nums, 0, end)
		}
		quickSort(nums)
	} else {
		merge := func(nums []int, i, mid, j int) {
			tmp := make([]int, j-i+1)
			// 合并区间 [i,mid] [mid+1,j]
			start, l, r := 0, i, mid+1
			for ; l <= mid && r <= j; start++ {
				if nums[l] <= nums[r] {
					tmp[start] = nums[l]
					l++
				} else {
					tmp[start] = nums[r]
					r++
				}
			}
			s, e := r, j
			if l <= mid {
				s, e = l, mid
			}
			for k := s; k <= e; k, start = k+1, start+1 {
				tmp[start] = nums[k]
			}

			for k, index := i, 0; k <= j && index < len(tmp); k, index = k+1, index+1 {
				nums[k] = tmp[index]
			}
		}
		var mergeSortPart func(nums []int, i, j int)
		mergeSortPart = func(nums []int, i, j int) {
			//递归终止条件
			if i >= j {
				return
			}
			mid := (i + j) >> 1
			mergeSortPart(nums, i, mid)
			mergeSortPart(nums, mid+1, j)
			merge(nums, i, mid, j)
		}
		// 归并
		mergeSort := func(nums []int) {
			l := len(nums)
			end := l - 1
			mergeSortPart(nums, 0, end)
		}
		mergeSort(nums)
	}
	return nums
}

// Maximum subarray
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	// 构建动态规划数组
	dp := make([]int, l)
	// 定义初解和最大值临时变量
	dp[0] = nums[0]
	m, max := math.MinInt, func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m = max(dp[0], m)
	for i := 1; i < l; i++ {
		// fmt.Println(dp)
		// 如果上一步是大于0,那就可以累加
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		m = max(dp[i], m)
	}
	return m
}

// Merge two sorted lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 哨兵节点
	head := &ListNode{}
	start := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			start.Next = list1
			list1 = list1.Next
		} else {
			start.Next = list2
			list2 = list2.Next
		}
		start = start.Next
	}
	t := list1
	if list1 == nil {
		t = list2
	}
	for t != nil {
		start.Next = t
		t = t.Next
		start = start.Next
	}
	return head.Next
}

// 2sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k := range nums {
		tmp := target - nums[k]
		if v, ok := m[tmp]; ok {
			return []int{v, k}
		}
		m[nums[k]] = k
	}
	return []int{}
}

// Binary tree level order traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// 使用BFS进行解答
	// 构造队列queue,用来调整层级数据
	bfs := func(queue []*TreeNode) {
		for len(queue) > 0 {
			level := []int{}
			// 记录操作前队列的长度,方便划分层级
			l := len(queue)
			// 遍历原队列长度
			for i := 0; i < l; i++ {
				// 拿出头节点
				node := queue[0]
				// 出队
				queue = queue[1:]
				level = append(level, node.Val)
				if node.Left != nil {
					// 入队
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					// 入队
					queue = append(queue, node.Right)
				}
			}
			// 遍历结束就是需要的结果
			res = append(res, level)
		}
	}
	bfs([]*TreeNode{root})
	return res
}

// Best time to buy and sell stock
func maxProfit(prices []int) int {
	l, maxP, k, max := len(prices), 0, 1, func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if l <= 2 {
		return 0
	}
	dp := make([][]int, k+1)
	for index := range dp {
		dp[index] = make([]int, l)
	}

	// 答案来源: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/39608/a-clean-dp-solution-which-generalizes-to-k-transactions
	// j 表示 prices[j]的价格
	// f[k, j] 代表在价格[j]之前的最大利润（注意：不是以价格[j]结束），最多使用k个交易
	// f[k, j] = max(f[k, j-1], prices[j] - prices[jj] + f[k-1, jj]) {0 <= jj <= j-1}
	//         = max(f[k, j-1], prices[j] + max(f[k-1, jj] - prices[jj]))
	// f[0, j] = 0; 0次交易就是0利润
	// f[k, 0] = 0; 如果价格保持不变,不管交易多少次收益都是0
	for i := 1; i <= k; i++ {
		tmaxp := dp[i-1][0] - prices[0]
		for j := 1; j < l; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+tmaxp)
			tmaxp = max(tmaxp, dp[i-1][j]-prices[j])
			maxP = max(maxP, dp[i][j])
		}
	}
	return maxP
}

func maxProfit2(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	// 构建结果存储
	// dp[i][0] 表示第i天不持有股票
	// dp[i][1] 表示第i天持有股票
	// 状态转移方程
	// dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
	// dp[i][1] = max(dp[i-1][1],-prices[i])
	dp, max := make([][]int, l), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for k := range dp {
		dp[k] = make([]int, 2)
	}
	// 构建初解
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	//遍历价格数组
	for i := 1; i < l; i++ {
		// 第i天不持有的最大收益应该是第i-1天不持有的最大收益和第i-1天持有股票,第i天卖出的最大收益之中的最大值
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天持有的最大收益应该是第i-1天持有的最大收益和第i-1天没有持有股票,第i天买入股票的最大收益之中的最大值
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	// 结果肯定是卖出的收益
	return dp[l-1][0]
}

// Linked list cycle
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 通过快慢指针,两个指针只要相遇,必有环
	// 注意: 快指针的快是每次走的快
	slow, fast := head, head.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// Number of islands
func numIslands(grid [][]byte) int {
	n, m, res := len(grid), len(grid[0]), 0
	// 定义深度优先函数
	var dfs func(grid [][]byte, i, j int) int
	dfs = func(grid [][]byte, i, j int) int {
		// 判断是不是越界或海水
		// 0 海水
		// 1 陆地
		// 2 已经遍历过的陆地
		if i >= n || j >= m || i < 0 || j < 0 || grid[i][j] != '1' {
			return 0
		}
		if grid[i][j] == '1' {
			// 标记已经遍历过了
			grid[i][j] = '2'
		}
		return 1 + dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)

	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num := dfs(grid, i, j)
			if num != 0 {
				res++
			}
		}
	}
	return res
}

// Search in rotated sorted array
func searchRotatedArray(nums []int, target int) int {
	l := len(nums)
	if l < 1 {
		return -1
	}
	// 初始化边界和循环退出条件
	for lo, hi := 0, l-1; lo <= hi; {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[lo] == target {
			return lo
		}
		if nums[hi] == target {
			return hi
		}
		if nums[lo] < nums[hi] {
			if target > nums[mid] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			// 判断有序
			if nums[lo] < nums[mid] {
				// 不在有序的就是在无序里根绝条件收敛
				if target > nums[mid] || target < nums[lo] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else {
				if target < nums[mid] || target > nums[hi] {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}
		}
	}
	return -1
}

// Binary tree zigzag level order traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 使用广度优先算法按层遍历
	// 根据遍历的层数奇偶来取值
	bfs := func(queue []*TreeNode) [][]int {
		var node *TreeNode
		res := [][]int{}
		start := 0
		// 使用队列来遍历节点数据
		for len(queue) > 0 {
			l := len(queue)
			tmpVal := []int{}
			// 取值
			if start%2 != 0 {
				for i := 0; i < l; i++ {
					tmpVal = append(tmpVal, queue[i].Val)
				}
			} else {
				for i := l - 1; i >= 0; i-- {
					tmpVal = append(tmpVal, queue[i].Val)
				}
			}

			// 处理下一次循环的队列
			for i := 0; i < l; i++ {
				node = queue[0]
				queue = queue[1:]
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
			}
			res = append(res, tmpVal)
			start++
		}
		return res
	}
	return bfs([]*TreeNode{root})
}

// Ontersection of two linked lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 使用hash表来记录 O(m+n) O(m)
	/*
		hashAuxiliary := func(headA, headB *ListNode) *ListNode {
			m := make(map[*ListNode]interface{})
			for headA != nil {
				m[headA] = nil
				headA = headA.Next
			}
			for headB != nil {
				if _, ok := m[headB]; ok {
					return headB
				}
				headB = headB.Next
			}
			return nil
		}
	*/
	// 采用跑圈原理,两个指针a b,a先跑A链再跑B链, b先跑B链再跑A链,如果二者能相遇,那就是重合的第一个节点相遇
	// a跑完 A+B链 或 b跑完B+A链则表示无重合
	a, b, flaga, flagb := headA, headB, false, false
	for a != b {
		a = a.Next
		if a == nil && !flaga {
			flaga = true
			a = headB
		}
		b = b.Next
		if b == nil && !flagb {
			flagb = true
			b = headA
		}
	}
	return a
}

// Valid parentheses
func isValid(s string) bool {
	if s == "" {
		return false
	}
	sArr, stack, record := []byte(s), []byte{}, map[byte]byte{
		// 构建正确的括号映射组
		40: 41, 91: 93, 123: 125,
	}
	for k := range sArr {
		// 栈里的数据作为前驱进行比较
		if len(stack) > 0 && record[stack[len(stack)-1]] == sArr[k] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, sArr[k])
	}
	return len(stack) == 0
}

// Lowest common ancestor of a binary tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 使用深度优先进行递归操作
	// 因为 p q 节点本身也可以是最近公共祖先
	// 从root开始中序遍历(遍历顺序无要求),可以分为以下几种情况
	// 1. 如果左右返回nil,表示该root下,p,q没有公共祖先
	// 2. 左右均不空,那root本身就是最近公共祖先
	// 3. 左不为空,右为空,左为最近公共祖先,左为空,右不为空,右为最近公共祖先
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == p || node == q || node == nil {
			return node
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		switch true {
		case left == right && right == nil:
			return nil
		case left == nil && right != nil:
			return right
		case left != nil && right == nil:
			return left
		case left != nil && right != nil:
			return node
		}
		return nil
	}
	return dfs(root)
}

// Longest palindromic substring
func longestPalindrome(s string) string {
	sArr := []byte(s)
	l := len(sArr)
	if l < 2 {
		return s
	}
	res, maxL := [2]int{}, math.MinInt
	// 中心位置寻找
	findFunc := func(sArr []byte, left, right int) [2]int {
		for left >= 0 && right < l {
			if sArr[left] == sArr[right] {
				left--
				right++
			} else {
				break
			}
		}
		return [2]int{left + 1, right - left - 1}
	}
	// 遍历数组
	for k := range sArr {
		odd := findFunc(sArr, k, k)
		even := findFunc(sArr, k, k+1)
		if odd[1] > even[1] {
			even = odd
		}
		if maxL < even[1] {
			res = even
			maxL = even[1]
		}
	}

	return s[res[0] : res[0]+res[1]]
}

// Merge sorted array
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// 因为都是有序数组,且 nums1 是放合并数组后的结果
	// 可以从大到小进行比较,一次放在 nums1 右边
	// 如果 nums1 有效长度遍历完,那就直接倒插入 nums2 剩余数值
	// 如果 nums2 有效长度遍历完,那么 nums1 就是已经合并的结果,直接 返回
	right, i, j := m+n-1, m-1, n-1
	for ; i >= 0 && j >= 0; right-- {
		if nums2[j] >= nums1[i] {
			nums1[right] = nums2[j]
			j--
		} else {
			nums1[i], nums1[right] = nums1[right], nums1[i]
			i--
		}
	}
	if i == -1 {
		for s := j; s >= 0; s-- {
			nums1[right] = nums2[s]
			right--
		}
	}
}

// Permutations
func permute(nums []int) [][]int {
	box := [][]int{}
	l := len(nums)
	if l == 0 {
		return box
	}
	var dfs func(depth int, used []bool, path []int)
	dfs = func(depth int, used []bool, path []int) {
		if depth == len(nums) {
			newTmp := make([]int, len(path))
			copy(newTmp, path)
			box = append(box, newTmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(depth+1, used, path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0, make([]bool, len(nums)), make([]int, 0, len(nums)))
	return box
}

// Merge k sorted lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	// 合并两个有序数组
	mergeTwoLists := func(list1 *ListNode, list2 *ListNode) *ListNode {
		head := &ListNode{}
		start := head
		for list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				start.Next = list2
				list2 = list2.Next
			} else {
				start.Next = list1
				list1 = list1.Next
			}
			start = start.Next
		}
		if list1 != nil {
			start.Next = list1
		} else {
			start.Next = list2
		}
		return head.Next
	}

	// 采用分治法
	var merge func(lists []*ListNode, left, right int) *ListNode
	merge = func(lists []*ListNode, left, right int) *ListNode {
		// 递归结束条件就是左右相等 || 左边大于右边
		if left == right {
			return lists[left]
		}
		if left > right {
			return nil
		}
		mid := (left + right) >> 1
		return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
	}

	return merge(lists, 0, l-1)
}

// Spiral matrix
func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	box := make([]int, 0, m*n)
	left, right, top, bottom := 0, m-1, 0, n-1
	for {
		// 向右走
		for walk := left; walk <= right; walk++ {
			box = append(box, matrix[top][walk])
		}
		top++
		if top > bottom {
			break
		}
		// 向下走
		for walk := top; walk <= bottom; walk++ {
			box = append(box, matrix[walk][right])
		}
		right--
		if left > right {
			break
		}
		// 向左走
		for walk := right; walk >= left; walk-- {
			box = append(box, matrix[bottom][walk])
		}
		bottom--
		if top > bottom {
			break
		}
		// 向上走
		for walk := bottom; walk >= top; walk-- {
			box = append(box, matrix[walk][left])
		}
		left++
		if left > right {
			break
		}
	}
	return box
}

// Reverse linked list ii
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var pre, start, end, next *ListNode
	c, s := 1, head
	for s != nil {
		if c+1 == left {
			pre, start = s, s.Next
		}
		if c == left {
			start = s
		}
		if c == right {
			end, next = s, s.Next
		}
		s = s.Next
		c++
	}
	if start == nil || end == nil {
		return head
	}
	end.Next = nil
	reverseList := func(s *ListNode) *ListNode {
		var pre *ListNode
		pre, cur := nil, s
		for cur != nil {
			tmp := cur.Next
			cur.Next = pre
			cur, pre = tmp, cur
		}
		return pre
	}
	if pre != nil {
		pre.Next = reverseList(start)
	} else {
		reverseList(start)
		head = end
	}
	start.Next = next
	return head
}

// Linked list cycle ii
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s, f := head, head
	// 定义环形链表环入口之前的节点数 a,环节点数 b
	// 找两个指针重合位置,此时
	// f(快指针,走两步)走了的节点数设为 x ,s(慢指针,走一步)走了的节点数设为 y
	// f比s多走一倍的节点数 x = 2y,因为环状链表, x = y + n*b (f比s多走了n圈环才能相遇的)
	// y = nb, x = 2*nb,重合时 s 走了 n 个环长, f 走了 2n 个环长
	// 从链表头走到环形入口的节点数是: a + n*b (环形入口后都是环,所以绕 n 都能回入口)
	// 此时 s,f 相遇,那么 f 再走 a 个节点则能到达入口
	// 利用一个新指针 h ,从链表头开始走 a 个节点则: h 与 f 相遇则为入口节点
	for {
		// 标识无环
		if f == nil || f.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next
		if f == s {
			break
		}
	}
	f = head
	for s != f {
		f, s = f.Next, s.Next
	}
	return f
}

// Add strings
func addStrings(num1 string, num2 string) string {
	res, i, j, up := "", len(num1)-1, len(num2)-1, 0
	// 利用两个字符串数字从后相加
	// 利用十进制进位来处理 十进制数 xyz, z / 10 标识是否进位 z % 10 标识当前位置 z 的值
	for i >= 0 || j >= 0 {
		var n1, n2 = 0, 0
		if i >= 0 {
			n1, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(string(num2[j]))
		}
		tmp := n1 + n2 + up
		up = tmp / 10
		res = strconv.Itoa(tmp%10) + res
		i--
		j--
	}
	if up != 0 {
		res = "1" + res
	}
	return res
}

// Longest increasing subsequence
func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	// 构建初始dp数组,基础值为1
	dp := make([]int, l)
	for k := range dp {
		dp[k] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := math.MinInt
	// 位置i的dp[i]最大值是 j表示[0,i)的值
	// 1. nums[i] > nums[j] 则i是所在位置j的严格后序,此时对应dp[j]++
	// 2. nums[i] <= nums[j] 不满足条件则跳出
	// dp[i] = max(dp[i],dp[j] + 1) j => [0,j)
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(dp[i], res)
	}
	return res
}

// Trapping rain water
func trap(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	// 接水我们以当前位置是否能积水为参考标准
	// 当前位置高为h,左边最高设为lm,右边最高设为rm
	// 只有当前位置高度h小于min(lm,rm)的时候才能有积水等于 min(lm,rm) - h
	// 建立左右指针 left right,首先通过比较 height[left] height[right]的高度确定接水的短板
	// 较短的一边指针位置与对应边界进行比较可以得出是否接住雨水
	left, right, lm, rm, sum := 0, len(height)-1, 0, 0, 0
	for left < right {
		// 接的雨水依赖右边
		if height[left] > height[right] {
			if rm > height[right] {
				sum += rm - height[right]
			} else {
				rm = height[right]
			}
			right--
		} else {
			if lm > height[left] {
				sum += lm - height[left]
			} else {
				lm = height[left]
			}
			left++
		}
	}
	return sum
}

// Binary tree maximum path sum
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	sum := math.MinInt
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// 当前节点具有的最大路径和(包括子树)
		tmpNum := left + right + node.Val
		// 与已知节点的路径和比较,选取最大
		sum = max(tmpNum, sum)
		// 输出的最大路径和,从当前根节点下来,只能走左或走右,不能走向左边后又走向右边
		return max(max(left, right)+node.Val, 0)
	}
	dfs(root)
	return sum
}

// Reorder list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// 1, 2, 3, 4, 5
	// 1, 5, 2, 4, 3
	// 构建一个数组存储链表节点
	s := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		s = append(s, node)
	}
	// 拿首尾节点的index
	l, r := 0, len(s)-1
	for l < r {
		s[l].Next = s[r]
		l++
		// 偶数的时候会提前结束
		if l == r {
			break
		}
		s[r].Next = s[l]
		r--
	}
	s[l].Next = nil
}

// Binary tree inorder traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var (
		ergodic func(node *TreeNode)
		res     []int
	)
	ergodic = func(node *TreeNode) {
		if node == nil {
			return
		}
		ergodic(node.Left)
		res = append(res, node.Val)
		ergodic(node.Right)
	}
	ergodic(root)
	return res
}

// Binary search
func search(nums []int, target int) int {
	low, hig := 0, len(nums)-1
	for low <= hig {
		mid := (low + hig) >> 1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			hig = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// Binary tree right side view
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res   []int
		queue = []*TreeNode{root}
	)
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == l-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

// Median of two sorted arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s1, l1, s2, l2 := 0, len(nums1), 0, len(nums2)
	l := l1 + l2
	if l == 0 {
		return 0
	}
	minLeft, minLeftV, minRight, minRightV, s, min := 0, -1, 0, -1, 0, l>>1
	if l%2 == 0 {
		minRight, minLeft = min, min-1
	} else {
		minRight, minLeft = min, min
	}

	for ; s1 < l1 && s2 < l2; s++ {
		val := 0
		if nums1[s1] < nums2[s2] {
			val = nums1[s1]
			s1++
		} else {
			val = nums2[s2]
			s2++
		}
		if s == minLeft {
			minLeftV = val
		}
		if s == minRight {
			minRightV = val
		}
	}
	if minLeftV == -1 || minRightV == -1 {
		var (
			i, j int
			num  []int
		)
		if s1 == l1 {
			i, j, num = s2, l2, nums2
		} else {
			i, j, num = s1, l1, nums1
		}
		for k := i; k < j; k++ {
			if s == minLeft {
				minLeftV = num[k]
			}
			if s == minRight {
				minRightV = num[k]
			}
			s++
		}
	}

	return float64(minLeftV+minRightV) / 2
}

// Merge intervals
func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res, slow, quick := [][]int{intervals[0]}, 0, 1
	for ; quick < l; quick++ {
		if res[slow][0] <= intervals[quick][0] && intervals[quick][0] <= res[slow][1] {
			var v []int
			if intervals[quick][1] > res[slow][1] {
				v = []int{res[slow][0], intervals[quick][1]}
			} else {
				v = []int{res[slow][0], res[slow][1]}
			}
			res[slow] = v
		} else {
			res = append(res, intervals[quick])
			slow++
		}
	}
	return res
}

// Remove nth node from end of list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := make([]*ListNode, 0, 500)
	for s := head; s != nil; s = s.Next {
		m = append(m, s)
	}
	l := len(m)
	index := l - n
	if l < 2 {
		return nil
	}
	switch index {
	case l - 1:
		m[index-1].Next = nil
	case 0:
		head = head.Next
	default:
		m[index-1].Next = m[index].Next
	}
	return head
}

// Climbing stairs
func climbStairs(n int) int {
	if n < 3 {
		s := [3]int{}
		s[0], s[1], s[2] = 0, 1, 2
		return s[n]
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Sort list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var normalFunc = func(head *ListNode) *ListNode {
		m := []*ListNode{}
		for i := head; i != nil; i = i.Next {
			m = append(m, i)
		}
		sort.Slice(m, func(i, j int) bool {
			return m[i].Val < m[j].Val
		})
		l := len(m) - 1
		for i := 0; i < l; i++ {
			m[i].Next = m[i+1]

		}
		m[l].Next = nil
		return m[0]
	}
	return normalFunc(head)
}

// Remove duplicates from sorted list ii
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dumy := &ListNode{Next: head}
	pre, cur, next := dumy, dumy.Next, dumy.Next.Next
	for next != nil {
		if cur.Val == next.Val {
			if next.Next == nil || next.Next.Val != cur.Val {
				pre.Next = next.Next
				if pre.Next != nil {
					cur = pre.Next
					next = pre.Next.Next
					continue
				} else {
					break
				}
			}
			next = next.Next
			continue
		} else {
			pre, cur, next = cur, next, next.Next
		}
	}
	return dumy.Next
}

// Sqrtx
// 牛顿迭代或者二分查找
func mySqrt(x int) int {
	if false {
		s := 1.0
		target := float64(x)
		for {
			tmp := (s + target/s) / 2.0
			if math.Abs(tmp-s) <= 0.1 {
				s = tmp
				break
			}
			s = tmp
		}
		return int(math.Floor(s))
	} else {
		ans, low, high := 0, 0, x
		for low <= high {
			mid := low + (high-low)/2
			if mid*mid <= x {
				ans = mid
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return ans
	}
}

// Next permutation
func nextPermutation(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	// 构建从后向前的快慢指针
	s, q, index := l-1, l-2, l-1
	// 向前遍历,寻找第一个相邻可交换的位置,跳出循环时[s,l)必然降序
	for q >= 0 && nums[q] >= nums[s] {
		s--
		q--
	}

	// 表示整个序列非升序,尝试在[s,l)之间找到一个比nums[q]小的数
	// 经过下面操作,此时[s,l)任然是降序
	if q >= 0 {
		for index >= s && nums[q] >= nums[index] {
			index--
		}
		if nums[q] < nums[index] {
			nums[q], nums[index] = nums[index], nums[q]
		}
	}

	// 对[s,l)进行升序排列
	for i, j := s, l-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Add two numbers
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, progression := &ListNode{}, 0
	s := head
	var s1, s2 *ListNode
	setFunc := func(a, b, c int, s *ListNode) (*ListNode, int) {
		tmp := a + b + c
		c = tmp / 10
		s.Next = &ListNode{Val: tmp % 10}
		s = s.Next
		return s, c
	}
	for s1, s2 = l1, l2; s1 != nil && s2 != nil; s1, s2 = s1.Next, s2.Next {
		s, progression = setFunc(s1.Val, s2.Val, progression, s)
	}

	endFunc := func(end *ListNode) {
		for ; end != nil; end = end.Next {
			s, progression = setFunc(0, end.Val, progression, s)
		}
		if progression == 1 {
			s.Next = &ListNode{Val: 1}
		}
	}

	if s1 == nil {
		endFunc(s2)
	} else {
		endFunc(s1)
	}
	return head.Next
}

// String to integer atoi
func myAtoi(s string) int {
	asBytes, res := []byte(s), int32(0)
	if len(asBytes) == 0 {
		return int(res)
	}
	index, l := 0, len(asBytes)
	// 去掉前置空格
	for index < l && asBytes[index] == ' ' {
		index++
	}
	// 防止极端情况全空格
	if index == l {
		return int(res)
	}
	// 判断是否有符号,首位有效
	sign := int32(1)
	switch asBytes[index] {
	case '+':
		index++
	case '-':
		index++
		sign = -1
	}
	// 开始循环后面得byte
	for ; index < l; index++ {
		currChar := asBytes[index]
		// 非数字字符直接跳过
		if currChar < '0' || currChar > '9' {
			break
		}
		// 拿上一步的结果来判断是否溢出
		last := res
		// 十进制计算 123 => 0*10+1 => 1*10+2 => 12*10+3 => 123
		res = res*10 + int32(currChar-'0')
		// 判断是否溢出
		if last != res/10 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}
	return int(res * sign)
}

// Longest common subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)

	if l1 == 0 || l2 == 0 {
		return 0
	}
	// 动态规划求解,首先定义状态数组 dp[i][j]
	// dp[i][j] 表示 text1[0,i-1] 与 text2[0,j-1] 之间的最长公共子序列数
	// dp[0][0] 表示空字符串,所以要考虑增加空字符串的可能性
	dp := make([][]int, l1+1)
	for k := range dp {
		dp[k] = make([]int, l2+1)
	}
	// 构建状态转移方程
	// 如果 text1[i-1] == text2[j-1],则表示末尾相同,那么 dp[i][j] = dp[i-1][j-1] + 1
	// 反之, dp[i][j] = max(dp[i][j-1], dp[i-1][j])
	b1, b2, max := []byte(text1), []byte(text2), func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if b1[i-1] == b2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}
