package array

import (
	"math"
	"sort"
)

// Sum of two numbers
func twoSum(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	for k, v := range nums {
		c := target - v
		if index, ok := m[c]; ok {
			return []int{index, k}
		}
		m[v] = k
	}
	return nil
}

// Find the median of two positive-order arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := l1 + l2
	lmax1, lmax2, rmin1, rmin2, c1, c2, lo, hi := 0, 0, 0, 0, 0, 0, 0, l1<<1
	for lo <= hi {
		c1 = (lo + hi) >> 1
		c2 = k - c1
		if c1 == 0 {
			lmax1 = math.MinInt
		} else {
			lmax1 = nums1[(c1-1)>>1]
		}
		if c1 == 2*l1 {
			rmin1 = math.MaxInt
		} else {
			rmin1 = nums1[c1>>1]
		}
		if c2 == 0 {
			lmax2 = math.MinInt
		} else {
			lmax2 = nums2[(c2-1)>>1]
		}
		if c2 == 2*l2 {
			rmin2 = math.MaxInt
		} else {
			rmin2 = nums2[c2>>1]
		}
		if lmax1 > rmin2 {
			hi = c1 - 1
		} else if lmax2 > rmin1 {
			lo = c1 + 1
		} else {
			break
		}
	}
	return (math.Max(float64(lmax1), float64(lmax2)) + math.Min(float64(rmin1), float64(rmin2))) / 2
}

// Containers for the most water
func maxArea(height []int) int {
	index := len(height) - 1
	l1, l2, area := 0, index, 0

	for l1 < l2 {
		l := l2 - l1
		// 面积取决于短板
		// 长板往内移动时遇到更长的板没有用,取决于最短的板,此时长度还缩短了
		// 短板往内移动时可能遇到比当前长的短板,有机会面积增大
		if height[l1] > height[l2] {
			area = max(area, height[l2]*l)
			l2--
		} else {
			area = max(area, height[l1]*l)
			l1++
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sum of three numbers
func threeSum(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	//特判,数组长度小于3就返回
	if l < 3 {
		return res
	}
	// 排序,从小到大
	sort.Ints(nums)
	// 遍历排序后的数组
	for i, _ := range nums {
		// 已经排好序,当nums[i]>0时.后面不会出现和为零了
		if nums[i] > 0 {
			return res
		}
		// 过滤重复的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo := i + 1
		hi := l - 1
		for lo < hi {
			r := nums[i] + nums[lo] + nums[hi]
			if r == 0 {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				// 过滤重复的值
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				// 过滤重复的值
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				hi--
				lo++
			} else if r > 0 { // 和大于0,那么右边界向左,总和缩小
				hi--
			} else { // 和小于0,那么左边界向右,总和增大
				lo++
			}
		}
	}
	return res
}

// Sum of the nearest three numbers
func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt
	l := len(nums)
	//特判,数组长度小于3就返回
	if l < 3 {
		return 0
	}
	// 排序,从小到大
	sort.Ints(nums)
	//定义处理
	approachNumber := func(a, b, c int) int {
		l := math.Abs(float64(a) - float64(c))
		r := math.Abs(float64(b) - float64(c))
		if l > r {
			return b
		}
		return a
	}
	// 遍历排序后的数组
	for i, _ := range nums {
		// 过滤重复的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo := i + 1
		hi := l - 1
		for lo < hi {
			r := nums[i] + nums[lo] + nums[hi]
			tmp := r - target
			if tmp == 0 {
				return r
			} else if r > target { // 和大于 target,那么右边界向左,总和缩小
				// 过滤重复的值
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				hi--
			} else { // 和小于 target,那么左边界向右,总和增大
				// 过滤重复的值
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
				lo++
			}
			res = approachNumber(res, r, target)
		}
	}
	return res
}

// Sum of four numbers
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	l := len(nums)
	//特判,数组长度小于4就返回
	if l < 4 {
		return res
	}
	// 排序,从小到大
	sort.Ints(nums)
	// 遍历排序后的数组
	for i := 0; i < l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 第一个数固定了,从小到大的顺序相加,只要和大于target,则后面的不可能出现等于target了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			return res
		}
		// 第一个数固定了,从大到小的顺序相加,只要和小于 target,第一个数和剩下的任意三个数相加都会小于,直接跳过
		if nums[i]+nums[l-3]+nums[l-2]+nums[l-1] < target {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 前两个数固定了,从小到大的顺序相加,只要和大于target,则后面的不可能出现等于target了
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 前两个数固定了,从大到小的顺序相加,只要和小于 target,和剩下的任意两个数相加都会小于,直接跳过
			if nums[i]+nums[j]+nums[l-2]+nums[l-1] < target {
				continue
			}
			lo := j + 1
			hi := l - 1
			for lo < hi {
				tmp := nums[i] + nums[j] + nums[lo] + nums[hi]
				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[lo], nums[hi]})
					//重复值跳过
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					//重复值跳过
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if tmp > target { // 和大于 target,那么右边界向左,总和缩小
					hi--
				} else { // 和小于 target,那么左边界向右,总和增大
					lo++
				}
			}
		}
	}
	return res
}

// Remove duplicate items from an ordered array
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	// 快慢指针,慢指针用来放不重复的数,快指针用来过滤重复
	i := 0
	for j := 1; j < len(nums); j++ {
		// 已经排好序了,快指针j遍历 1 ~ len(nums)-1 位置
		// 快指针 j 与 j-1 不等表示当前j位置的数可以放到慢指针i+1的位置
		if nums[j] != nums[j-1] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

// Remove elements
func removeElement(nums []int, val int) int {
	// l := 0 // 定义覆盖位置指针l
	// for _,v := range nums {
	//     if v != val { // 不等则将指针l位置覆盖,同时移动l向右
	//         nums[l] = v
	//         l++
	//     } // 相等子继续遍历,指针l位置不变
	// }
	// return l
	lo, hi := 0, len(nums) // 定义左右指针
	for lo < hi {
		if nums[lo] == val { // 相等表示可以覆盖左指针位置,然后右指针减一,左指针不变,下轮继续判断
			nums[lo] = nums[hi-1]
			hi--
		} else { // 不等则移动左指针,进行下轮判断
			lo++
		}
	}
	return lo
}

// Next permutation
func nextPermutation(nums []int) {
	// 翻译下 给定数字序列的字典序中下一个更大的排列。如果不存在下一个更大的排列，则将数字重新排列成最小的排列
	// 等同于 给定若干个数字将其组合为一个整数.如何将这些数字重新排列,以得到下一个更大的整数.如 123 下一个更大的数为 132.如果没有更大的整数,则输出最小的整数
	l := len(nums)
	if l < 2 {
		return
	}
	// 从后往前开始匹配,设置两个指针i,j
	// 找到第一个 nums[i] < nums[j] 处
	i, j, k := l-2, l-1, l-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 在在[j,l)中寻找 nums[i] < nums[k] 满足的指针k的位置
	// 调换 i,k 位置
	// i<0标识数组元素是降序的,只要变为升序即可
	if i >= 0 {
		for k >= j && nums[i] >= nums[k] {
			k--
		}
		if nums[i] < nums[k] {
			nums[i], nums[k] = nums[k], nums[i]
		}
	}

	// [j,l)转换为升序
	for i, j := j, l-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Search in rotated sorted array
func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 初始化边界和循环退出条件
	for lo, hi := 0, l-1; lo <= hi; {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return mid
		}
		// 数组数据二分,不包含旋转点的一定是正常升序
		// 目标在正常升序中,继续二分即可
		// 目标在有旋转点的部分,可以继续二分,重新判断按照上述逻辑判断
		if nums[0] <= nums[mid] { // 表示[0,mid)正常升序的,不包含旋转点,常规二分
			if nums[0] <= target && target < nums[mid] {
				hi = mid - 1
			} else { // 重新定义,mid值随之更新
				lo = mid + 1
			}
		} else { // [0,mid)包含旋转点,那么[mid,l)为正常升序,常规二分
			if nums[mid] < target && target <= nums[l-1] {
				lo = mid + 1
			} else { // 重新定义,mid值随之更新
				hi = mid - 1
			}
		}
	}
	return -1
}

// Find first and last position of element in sorted array
func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	lo, hi, L, R := 0, l-1, 0, 0
	// 定义[lo,mid],[mid+1,hi]区间,lo = mid+1 | hi = mid
	// target <= nums[mid] hi不断向左边界逼近,最终值为最左边第一个出现的target位置
	for lo < hi {
		mid := (lo + hi) >> 1
		if target <= nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[hi] != target {
		return []int{-1, -1}
	}
	L = hi
	lo, hi = 0, l-1
	// 定义[lo,mid-1],[mid,hi]区间,lo = mid | hi = mid-1
	// target >= nums[mid] lo不断向右边界逼近,最终值为最右边最后一个出现的target位置
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		if target >= nums[mid] {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	R = lo
	return []int{L, R}
}

// Search insert position
func searchInsert(nums []int, target int) int {
	l := len(nums)
	end := l - 1
	if l == 0 {
		return 0
	}
	if l == 1 {
		if target <= nums[0] {
			return 0
		} else if target > nums[0] {
			return 1
		}
	}
	// 单调且不重复可以直接判断头尾
	if target <= nums[0] {
		return 0
	}
	if target > nums[end] {
		return l
	} else if target == nums[end] {
		return end
	}
	lo, hi := 0, end
	// 区域分成[lo,mid],[mid+1,hi] lo = mid+1 | hi = mid
	for lo < hi {
		mid := (lo + hi) >> 1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	// lo , hi 在找数组中不在的元素都是收缩边界,可以任选
	return hi
}

// Valid sudoku
func isValidSudoku(board [][]byte) bool {
	var (
		// 数字 1~9 映射到数组[]int{}的位置就是 0~8
		//  l[2][3] = true 表示数字 4 在第 2 行已经出现过
		//  c[2][3] = true 表示数字 4 在第 2 列已经出现过
		//  box[0][0][3] = true 表示数字 4 在(0,0)九宫格中已经出现过
		l, c [9][9]bool    // 存储横,纵向数据
		box  [3][3][9]bool // 存储九宫格内数据
	)
	// 从左到右从上到下的遍历 O(n²)
	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			index := v - '1'
			if l[i][index] {
				return false
			}

			if c[j][index] {
				return false
			}

			if box[i/3][j/3][index] {
				return false
			}
			l[i][index] = true
			c[j][index] = true
			box[i/3][j/3][index] = true
		}
	}
	return true
}

// Sudoku solver
func solveSudoku(board [][]byte) {
	var (
		// 数字 1~9 映射到数组[]int{}的位置就是 0~8
		//  l[2][3] = true 表示数字 4 在第 2 行已经出现过
		//  c[2][3] = true 表示数字 4 在第 2 列已经出现过
		//  box[0][0][3] = true 表示数字 4 在(0,0)九宫格中已经出现过
		l, c [9][9]bool
		box  [3][3][9]bool
		tmp  [][2]int
		dfs  func(int) bool
	)

	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				tmp = append(tmp, [2]int{i, j})
			} else {
				index := v - '1'
				l[i][index] = true
				c[j][index] = true
				box[i/3][j/3][index] = true
			}
		}
	}

	dfs = func(p int) bool {
		if p == len(tmp) {
			return true
		}
		i, j := tmp[p][0], tmp[p][1]
		for v := byte(0); v < 9; v++ {
			if !l[i][v] && !c[j][v] && !box[i/3][j/3][v] {
				l[i][v] = true
				c[j][v] = true
				box[i/3][j/3][v] = true
				board[i][j] = v + '1'
				if dfs(p + 1) {
					return true
				}
				l[i][v] = false
				c[j][v] = false
				box[i/3][j/3][v] = false
			}
		}
		return false
	}
	dfs(0)
}

// Combination sum
func combinationSum(candidates []int, target int) [][]int {
	l := len(candidates)
	boxs := [][]int{}
	if l <= 0 {
		return boxs
	}
	// 排序,方便减枝
	sort.Ints(candidates)
	// 采用 选择 | 不选择 来构建递归树
	// var dfs func(target, idx int, tmp []int)
	// dfs = func(target, idx int, tmp []int) {
	// 	if target <= 0 || idx == l {
	// 		if target == 0 {
	// 			newTmp := make([]int, len(tmp))
	// 			copy(newTmp, tmp)
	// 			boxs = append(boxs, newTmp)
	// 		}
	// 		return
	// 	}

	// 	// 不选择
	// 	dfs(target, idx+1, tmp)
	// 	// 选择
	// 	v := target - candidates[idx]
	// 	if v >= 0 {
	// 		tmp = append(tmp, candidates[idx])
	// 		dfs(v, idx, tmp)
	// 		tmp = tmp[:len(tmp)-1]
	// 	}
	// }
	// dfs(target, 0, []int{})
	// 采用目标值减值构建递归树
	var dfs func(target, start int, tmp []int)
	dfs = func(target, start int, tmp []int) {
		if target <= 0 {
			if target == 0 {
				newTmp := make([]int, len(tmp))
				copy(newTmp, tmp)
				boxs = append(boxs, newTmp)
			}
			return
		}
		// 这里保证同层第二个节点不使用第一个节点在 candidates 里完全使用过的值
		// candidates = [2,3,5]
		// 第一个节点是(-2 -3 -5) 第二个节点(-3 -5)
		for i := start; i < l; i++ {
			v := target - candidates[i]
			if v < 0 {
				break
			}
			tmp = append(tmp, candidates[i])
			dfs(v, i, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(target, 0, []int{})
	return boxs
}

// Combination sum ii
func combinationSum2(candidates []int, target int) [][]int {
	box := [][]int{}
	l := len(candidates)
	if l <= 0 {
		return box
	}
	// 排序,加速剪枝判断
	sort.Ints(candidates)
	// 采用 选择 | 不选择 来构建递归树
	// 构建hash, [pos][0] [pos][1] 表示 candidates 某个数的值和次数
	// var m [][2]int
	// for _, n := range candidates {
	// 	if m == nil || n != m[len(m)-1][0] {
	// 		m = append(m, [2]int{n, 1})
	// 	} else {
	// 		m[len(m)-1][1]++
	// 	}
	// }
	// min := func(a, b int) int {
	// 	if a > b {
	// 		return b
	// 	}
	// 	return a
	// }
	// var dfs func(target, idx int, tmp []int)
	// dfs = func(target, idx int, tmp []int) {
	// 	if target <= 0 || idx == l || target < m[idx][0] {
	// 		if target == 0 {
	// 			newTmp := make([]int, len(tmp))
	// 			copy(newTmp, tmp)
	// 			box = append(box, newTmp)
	// 		}
	// 		return
	// 	}

	// 	dfs(target, idx+1, tmp) // 不选择

	// 	repeat := min(target/m[idx][0], m[idx][1]) // 通过用 目标值/当前位置的值 和 次数 判断,选取小的重复递归

	// 	for i := 1; i <= repeat; i++ {// 选择
	// 		tmp = append(tmp, m[idx][0])
	// 		dfs(target-i*m[idx][0], idx+1, tmp)
	// 	}
	// 	tmp = tmp[:len(tmp)-repeat]
	// }
	// dfs(target, 0, []int{})
	var dfs func(target, start int, tmp []int)
	dfs = func(target, start int, tmp []int) {
		if target <= 0 {
			if target == 0 {
				newTmp := make([]int, len(tmp))
				copy(newTmp, tmp)
				box = append(box, newTmp)
			}
			return
		}
		for i := start; i < l; i++ {
			v := target - candidates[i]
			if v < 0 {
				break
			}
			// 同层相同必须跳过,上下层相同可以
			if start < i && candidates[i] == candidates[i-1] {
				continue
			}
			tmp = append(tmp, candidates[i])
			// 已经使用过的不在使用
			dfs(v, i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(target, 0, []int{})
	return box
}
