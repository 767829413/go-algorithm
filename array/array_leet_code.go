package array

import (
	"math"
	"sort"
	"strings"
	"time"
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
	for i := range nums {
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
	for i := range nums {
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
func removeDuplicatesI(nums []int) int {
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
func searchI(nums []int, target int) int {
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

// First missing positive
func firstMissingPositive(nums []int) int {
	l := len(nums)
	// 建立hash表,存储所有nums 大于0的值
	m := make(map[int]int, l)
	for _, v := range nums {
		if v <= 0 {
			continue
		}
		m[v] = 1
	}
	// 从1到l进行遍历,因为找的是最小的未出现的数,转换为hash表里未出现的最小数即可
	// 为啥终止条件是l呢? 如果是 [1~n](顺序+1递增)那么nums长为n,最终最小值也就为n+1了
	// 如果是 [x~n](非顺序+1递增) 那么 1到n遍历最小的值肯定在1到l之间
	for i := 1; i <= l; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return l + 1
}

// Trapping rain water
func trap(height []int) int {
	l := len(height)
	sum := 0
	if l <= 0 {
		return sum
	}

	// 接水我们以当前位置是否能积水为参考标准
	// 当前位置高为h,左边最高设为lm,右边最高设为rm
	// 只有当前位置高度h小于min(lm,rm)的时候才能有积水等于 min(lm,rm) - h
	// 建立左右指针 lo hi,通过比较 height[lo] , height[hi]的高度确定当前位置依赖哪边高度
	lo, hi, lm, rm := 0, l-1, 0, 0
	for lo < hi {
		// 左边高度为依赖
		if height[lo] < height[hi] {
			// 判断lo位置是不是左边最高的
			if height[lo] >= lm {
				lm = height[lo]
			} else {
				// 不是最高则是可积水位置
				sum += lm - height[lo]
			}
			lo++
		} else {
			if height[hi] >= rm {
				rm = height[hi]
			} else {
				sum += rm - height[hi]
			}
			hi--
		}
	}
	return sum
}

// Jump game ii
func jump(nums []int) int {
	step, jmax, l, f := 0, 0, len(nums), 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 正向寻找,将每次的前进距离选取最大值
	// 最后一个位置跳过,因为要么超过最后位置,要么等于,都表示结束
	for i := 0; i < l-1; i++ {
		// 每个位置能到达的最远位置
		jmax = max(jmax, i+nums[i])
		// 到达最远位置边界
		if i == f {
			// fmt.Println("边界位置", f)
			// 更新最新能到达的最远位置
			f = jmax
			// 此时因为已经到了上一次最远的边界,将跳的次数更新
			step++
		}
		// fmt.Println(i, jmax, f)
	}
	return step
}

// Permutations
func permute(nums []int) [][]int {
	box := [][]int{}
	l := len(nums)
	if l == 0 {
		return box
	}
	var dfs func(depth int, used []bool, path []int)
	// 利用深度进行逐层循环遍历
	dfs = func(depth int, used []bool, path []int) {
		// 设置终止条件,深度
		if depth == l {
			// 必须用新的值去拷贝,不然数据会被重复设置
			// 切片底层还是数组,引用导致的
			newPath := make([]int, len(path))
			copy(newPath, path)
			box = append(box, newPath)
			return
		}
		for i := 0; i < l; i++ {
			// 剔除重复使用
			if used[i] {
				continue
			}
			// 开始尝试状态
			path = append(path, nums[i])
			used[i] = true
			dfs(depth+1, used, path)
			// 恢复到尝试前的状态
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0, make([]bool, l), []int{})
	return box
}

// Permutations ii
func permuteUnique(nums []int) [][]int {
	box := [][]int{}
	l := len(nums)
	if l == 0 {
		return box
	}
	// 排序,方便剪枝
	sort.Ints(nums)
	var dfs func(depth int, path []int, used []bool)
	dfs = func(depth int, path []int, used []bool) {
		if depth == l {
			// fmt.Println(path)
			newPath := make([]int, len(path))
			copy(newPath, path)
			box = append(box, newPath)
			return
		}
		for i := 0; i < l; i++ {
			// 同深度去重的关键是上一层同值不能被使用,被使用了当前同值不可跳过
			// 剩下的就是普通的回溯了
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(depth+1, path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{}, make([]bool, l))
	return box
}

// Rotate image
func rotate(matrix [][]int) {
	n := len(matrix)
	// 关键公式 90°旋转表示 matrix[row][col] <=> matrix[col][n-row-1]
	// 水平翻转 matrix[row][col] <=> matrix[n-row-1][col]
	// 主对角线翻转 matrix[n-row-1][col] <=> matrix[col][n-row-1]

	// 水平轴翻转 记住水平轴翻转是整体的一半
	for i := 0; i < (n >> 1); i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}

	// fmt.Println(matrix)
	// 对角线翻转 记住翻转的时候j<i
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// fmt.Println(matrix)
}

// Group anagrams
func groupAnagrams(strs []string) [][]string {
	// sort.Slice()
	box := make([][]string, 0, len(strs))
	// 字符串字典排序后作为hash key
	// m := make(map[string][]string)
	// for _, v := range strs {
	// 	tmp := []byte(v)
	// 	sort.Slice(tmp, func(i, j int) bool { return tmp[j] > tmp[i] })
	// 	index := string(tmp)
	// 	m[index] = append(m[index], v)
	// }
	// for _, v := range m {
	// 	box = append(box, v)
	// }
	// 使用计数,字母异位词之间包含的字母必定个数相同
	c := make(map[[26]int][]string)
	for _, v := range strs {
		tmp := [26]int{}
		for _, b := range v {
			tmp[b-'a']++
		}
		c[tmp] = append(c[tmp], v)
	}
	for _, v := range c {
		box = append(box, v)
	}
	return box
}

// N queens
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{{""}}
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}
	box := [][]int{}
	// 每一列的使用
	col := make([]bool, n)
	// 主对角线的使用 row(行) col(列) 关系是 row - col => [-3~3],总体+3,[0~6]
	main := make([]bool, 2*n-1)
	// 副对角线的使用 row(行) col(列) 关系是 row + col => [0~6]
	vice := make([]bool, 2*n-1)

	var dfs func(row int, path []int)
	dfs = func(row int, path []int) {
		if row == n {
			newPath := make([]int, len(path))
			copy(newPath, path)
			box = append(box, newPath)
			return
		}
		for i := 0; i < n; i++ {
			if col[i] || main[row-i+n-1] || vice[row+i] {
				continue
			}
			path = append(path, i)
			// 记录使用状态
			col[i] = true
			main[row-i+n-1] = true
			vice[row+i] = true

			dfs(row+1, path)

			// 状态还原
			col[i] = false
			main[row-i+n-1] = false
			vice[row+i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	ans := make([][]string, len(box))
	// 将获取的数组转化为字符串
	for k, v := range box {
		ans[k] = []string{}
		for _, index := range v {
			tmp := []byte{}
			for j := 0; j < n; j++ {
				if j == index {
					tmp = append(tmp, 'Q')
				} else {
					tmp = append(tmp, '.')
				}
			}
			ans[k] = append(ans[k], string(tmp))
		}
	}
	return ans
}

// Maximum subarray
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	// 动态规划经典题
	// 定义子问题 以当前nums[i]位置结尾的最大值
	// 描叙子问题之间联系 nums[i]位置的最大值是取决于nums[i-1]位置的最大值
	// 初值设置 nums[0]位置的最大值就是nums[0]
	dp := make([]int, l)
	dp[0] = nums[0]
	m := math.MinInt
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		m = max(dp[i], m)
	}
	m = max(dp[0], m)
	return m
}

// Spiral matrix
func spiralOrder(matrix [][]int) []int {
	box := []int{}
	l := len(matrix)
	if l == 0 {
		return box
	}
	// 构建上下,左右边界
	rowU, rowrD, colL, colR := 0, l-1, 0, len(matrix[0])-1
	// 向右走,向下走,向左走,向上走
	for {
		// 循环退出条件
		if colL > colR {
			break
		}

		for j := colL; j <= colR; j++ {
			box = append(box, matrix[rowU][j])
		}
		rowU++

		// 循环退出条件
		if rowU > rowrD {
			break
		}

		for i := rowU; i <= rowrD; i++ {
			box = append(box, matrix[i][colR])
		}
		colR--

		// 循环退出条件
		if colL > colR {
			break
		}

		for j := colR; j >= colL; j-- {
			box = append(box, matrix[rowrD][j])
		}
		rowrD--

		// 循环退出条件
		if rowU > rowrD {
			break
		}

		for i := rowrD; i >= rowU; i-- {
			box = append(box, matrix[i][colL])
		}
		colL++
	}
	return box
}

// Jump game
func canJump(nums []int) bool {
	l := len(nums)
	if l != 1 && nums[0] == 0 {
		return false
	}
	if l == 1 || nums[0] >= l-1 {
		return true
	}

	// 反向找可以跳的位置, O(n²)
	// end, find := l-1, false
	// for {
	// 	find = false
	// 	for i := 0; i < end; i++ {
	// 		if i+nums[i] >= end {
	// 			end = i
	// 			find = true
	// 		}
	// 	}
	// 	if end == 0 {
	// 		return true
	// 	}
	// 	// fmt.Println(find, end)
	// 	if !find {
	// 		return false
	// 	}
	// }

	// 维持连续正向跳最远 O(n)
	// 1. 如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
	// 2. 可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
	// 3. 如果可以一直跳到最后，就成功了
	index, max := 0, func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < l; i++ {
		// 当前位置已经大于前面最大能跳的位置了那必然是无法继续跳下去了
		if i > index {
			return false
		}
		index = max(index, nums[i]+i)
	}
	return true
}

// Merge intervals
func merge(intervals [][]int) [][]int {
	box := [][]int{}
	l := len(intervals)
	if l == 1 {
		return intervals
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 首先对每个区间的起始位置进行排序
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 构建一个对比的序列
	box = append(box, intervals[0])
	for i := 1; i < l; i++ {
		j := len(box) - 1
		// 如果a,b区间能合并,必然是 a区间[i,j]中的j >= b区间[x,y]中x
		if intervals[i][0] <= box[j][1] {
			box[j][1] = max(intervals[i][1], box[j][1])
		} else { // 不能合并那就得把当前值push到对比序列里
			box = append(box, intervals[i])
		}
	}
	return box
}

// Insert interval
func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	if l == 0 {
		return [][]int{newInterval}
	}
	intervals = append(intervals, newInterval)
	// 直接复用上一题的解答就ok了
	return merge(intervals)
}

// Spiral matrix ii
func generateMatrix(n int) [][]int {
	end := n * n
	if end == 1 {
		return [][]int{{1}}
	}
	box := make([][]int, n)
	for k := range box {
		box[k] = make([]int, n)
	}
	// 还是很简单的思路 向右走,向下走,向左走,向上走,控制好边界
	// 退出条件就是 左边界 > 右边界 上边界 > 下边界
	start, colL, colR, rowU, rowD := 1, 0, n-1, 0, n-1
	for {
		// 向右走
		for j := colL; j <= colR; j++ {
			box[rowU][j] = start
			start++
		}
		// 向右走完,上边界下移
		rowU++

		if rowU > rowD {
			break
		}
		// 向下走
		for i := rowU; i <= rowD; i++ {
			box[i][colR] = start
			start++
		}
		// 向下走完,右边界左移
		colR--

		if colL > colR {
			break
		}
		// 向左走
		for j := colR; j >= colL; j-- {
			box[rowD][j] = start
			start++
		}
		// 向左走完,下边界上移
		rowD--
		// 向上走
		for i := rowD; i >= rowU; i-- {
			box[i][colL] = start
			start++
		}
		// 向上走完,左边界右移
		colL++

		if colL > colR {
			break
		}
	}
	return box
}

// Unique paths ii
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	if n == 0 || m == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	// 动态规划求解,先推到子问题到最终大问题的求解递推式
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, m)
	}
	// 最初解确定
	dp[0][0] = 1
	// 最左侧可达解,只要有障碍物那后续都不可达,可达只有一条路径
	for i := 1; i < n && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	// 最上侧可达解,同上
	for j := 1; j < m && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			// 当前位置是障碍物的,那必然不可达
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 按递推公式进行递推求解
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

// Minimum path sum
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	if n == 0 || m == 0 {
		return 0
	}
	if m == n && n == 1 {
		return grid[0][0]
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 动态规划求解,先推到子问题到最终大问题的求解递推式
	// dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, m)
	}
	// 最初解确定
	dp[0][0] = grid[0][0]
	// 向下走的初始最小解
	for i := 1; i < n; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	// 向右走的初始最小解
	for j := 1; j < m; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			// 按递推公式进行递推求解
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[n-1][m-1]
}

// Plus one
func plusOne(digits []int) []int {
	l := len(digits)
	// 取最后一位数开始
	for i := l - 1; i >= 0; i-- {
		digits[i]++
		// 取余
		digits[i] %= 10
		// 取余不为0则返回
		if digits[i] != 0 {
			return digits
		}
	}
	// 到这就是全体进一位,扩充切片
	return append([]int{1}, digits...)
}

// Text justification
func fullJustify(words []string, maxWidth int) []string {
	l := len(words)
	box := []string{}
	var builder strings.Builder
	if l == 0 {
		for i := 0; i < maxWidth; i++ {
			builder.WriteString(" ")
		}
		return []string{builder.String()}
	}

	getP := func(nsn, indexN int) int {
		var p, p1, p2 int
		p1 = nsn % indexN
		p2 = nsn / indexN
		if p1 == 0 {
			p = p2
		} else {
			p = p2 + 1
		}
		return p
	}

	setSpace := func(nsn int, res []string) string {
		end := len(res) - 1
		indexN := end
		p := getP(nsn, indexN)
		newR := []string{}
		for k := range res {
			newR = append(newR, res[k])
			if k == end {
				break
			}
			if k == end-1 {
				for i := 0; i < nsn; i++ {
					newR = append(newR, " ")
				}
			} else {
				for i := 0; i < p; i++ {
					newR = append(newR, " ")
				}
				indexN--
				nsn -= p
				p = getP(nsn, indexN)
			}

		}
		return strings.Join(newR, "")
	}
	setEndSpace := func(nsn int, res []string) string {
		end := len(res) - 1
		indexN := end
		if indexN == 0 {
			for i := 0; i < nsn; i++ {
				res = append(res, " ")
			}
			return strings.Join(res, "")
		} else {
			tmp := []string{}
			for k := range res {
				tmp = append(tmp, res[k])
				if k == end {
					for i := 0; i < nsn; i++ {
						tmp = append(tmp, " ")
					}
					break
				} else {
					tmp = append(tmp, " ")
					nsn--
				}
			}
			return strings.Join(tmp, "")
		}
	}
	type tmp struct {
		len     int
		realLen int
		str     []string
	}
	t := &tmp{
		len: 0,
		str: []string{},
	}
	for i := 0; i < l; i++ {
		sn := maxWidth - t.realLen
		flag := maxWidth - t.len - len(words[i])

		if flag <= 0 {
			if flag == 0 {
				t.str = append(t.str, words[i])
				t.realLen = t.realLen + len(words[i])
				t.len = t.len + len(words[i]) + 1
				sn = maxWidth - t.realLen
			}
			needStr := ""
			if len(t.str) == 1 {
				needStr = setEndSpace(sn, t.str)
			} else {
				needStr = setSpace(sn, t.str)
			}
			box = append(box, needStr)
			t = &tmp{
				len:     0,
				realLen: 0,
				str:     []string{},
			}
		}

		if i == l-1 && flag != 0 {
			t.str = append(t.str, words[i])
			t.realLen = t.realLen + len(words[i])
			t.len = t.len + len(words[i]) + 1
			sn = maxWidth - t.realLen
			needStr := setEndSpace(sn, t.str)
			box = append(box, needStr)
			return box
		}

		if flag != 0 {
			t.str = append(t.str, words[i])
			t.realLen = t.realLen + len(words[i])
			t.len = t.len + len(words[i]) + 1
		}

	}
	return box
}

// Set matrix zeroes
func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	if n == 0 || m == 0 {
		return
	}
	// 判断行列是否有0标志
	col0Flag, row0Flag := false, false
	// 判断行是否有0
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			row0Flag = true
			break
		}
	}

	// 判断列是否为0
	for j := 0; j < m; j++ {
		if matrix[0][j] == 0 {
			col0Flag = true
			break
		}
	}

	// 然后利用第一行第一列作为记录来标记
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}

	// 根据第一行第一列的标记开始置0
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 最后根绝条件将第一行第一列置0
	if row0Flag {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

	if col0Flag {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
}

// Search a 2d matrix
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	if n == 0 || m == 0 || target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}
	// 二分查找
	secondPartSearch := func(a []int, target int) bool {
		lo, hi := 0, len(a)-1
		if lo == hi {
			return a[lo] == target
		}
		// 定义[lo,mid],[mid+1,hi]区间,lo = mid+1 | hi = mid
		for lo <= hi {
			mid := (lo + hi) >> 1
			if target < a[mid] {
				hi = mid - 1
			} else if target > a[mid] {
				lo = mid + 1
			} else {
				return true
			}
		}
		return false
	}

	// 对第一列进行二分查找一个稍微大一点的值
	row, rowL, rowH := -1, 0, n-1
	for rowL <= rowH {
		mid := (rowL + rowH) >> 1
		if target <= matrix[mid][m-1] {
			row = mid
			rowH = mid - 1
		} else if target > matrix[mid][m-1] {
			rowL = mid + 1
		}
	}
	if row == -1 {
		return false
	}

	return secondPartSearch(matrix[row], target)
}

// Sort colors
func sortColors(nums []int) {
	// 整数 0、 1 和 2 分别表示红色、白色和蓝色
	// 按照红色、白色、蓝色顺序排列
	start, zero, two := 0, 0, len(nums)-1

	// 构建区间 [0,zero) [zero,two] (two,len(nums)-1]
	for start <= two {
		if nums[start] == 0 {
			nums[start], nums[zero] = nums[zero], nums[start]
			zero++
			start++
		} else if nums[start] == 1 {
			start++
		} else {
			nums[start], nums[two] = nums[two], nums[start]
			two--
		}
		// fmt.Println(start, zero, two, nums)
	}
}

// Subsets
func subsets(nums []int) [][]int {
	// 首先画出递归参照图
	l := len(nums)
	if l == 0 {
		return append([][]int{}, []int{})
	}
	if l == 1 {
		return append([][]int{{}}, nums)
	}
	box := [][]int{}
	var dfs func(start int, tmp []int)
	dfs = func(start int, tmp []int) {
		newTmp := make([]int, len(tmp))
		copy(newTmp, tmp)
		box = append(box, newTmp)
		for i := start; i < l; i++ {
			tmp = append(tmp, nums[i])
			// 当前递归 return 时候,应该是i+1而不是 start+1 组合排列是有区别的
			dfs(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, []int{})
	return box
}

// Word search
func exist(board [][]byte, word string) bool {
	// {'A', 'B', 'C', 'E'},
	// {'S', 'F', 'C', 'S'},
	// {'A', 'D', 'E', 'E'},
	n := len(board)
	m := len(board[0])
	wordArr := []byte(word)
	// 构建map记录已走
	mr := make([][]int, n)
	for k := range mr {
		mr[k] = make([]int, m)
	}
	flag := make(chan bool)
	var dfs func(depth, row, col int, wordArr []byte)
	dfs = func(depth, row, col int, wordArr []byte) {
		// 标识存在
		if depth == len(wordArr)-1 {
			flag <- true
			return
		}

		// 向下走
		if row+1 < n && wordArr[depth+1] == board[row+1][col] && mr[row+1][col] == 0 {
			// fmt.Println("向下走", string(board[row+1][col]), string(wordArr[depth+1]), depth+1, mr[row+1][col])
			mr[row+1][col] = 1
			dfs(depth+1, row+1, col, wordArr)
			mr[row+1][col] = 0
		}

		// 向右走
		if col+1 < m && wordArr[depth+1] == board[row][col+1] && mr[row][col+1] == 0 {
			// fmt.Println("向右走", string(board[row][col+1]), string(wordArr[depth+1]), depth+1, mr[row][col+1])
			mr[row][col+1] = 1
			dfs(depth+1, row, col+1, wordArr)
			mr[row][col+1] = 0
		}

		// 向左走
		if col >= 1 && wordArr[depth+1] == board[row][col-1] && mr[row][col-1] == 0 {
			// fmt.Println("向左走", string(board[row][col-1]), string(wordArr[depth+1]), depth+1, mr[row][col-1])
			mr[row][col-1] = 1
			dfs(depth+1, row, col-1, wordArr)
			mr[row][col-1] = 0
		}

		// 向上走
		if row >= 1 && wordArr[depth+1] == board[row-1][col] && mr[row-1][col] == 0 {
			// fmt.Println("向上走", string(board[row-1][col]), string(wordArr[depth+1]), depth+1, mr[row-1][col])
			mr[row-1][col] = 1
			dfs(depth+1, row-1, col, wordArr)
			mr[row-1][col] = 0
		}
	}

	go func() {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if wordArr[0] != board[i][j] {
					continue
				}
				mr[i][j] = 1
				dfs(0, i, j, wordArr)
				mr[i][j] = 0
			}
		}
		flag <- false
	}()

	select {
	case res := <-flag:
		return res
	case <-time.After(500 * time.Second):
		return false
	}
}

// Remove duplicates from sorted array ii
func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			if j < len(nums) && j > i+1 {
				nums = append(nums[:j], nums[j+1:]...)
			} else {
				j++
			}
		} else {
			i = j
			j += 1
		}
	}
	return j
}

// Search in rotated sorted array ii
func search(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	if l == 1 {
		return nums[0] == target
	}
	for low, hig := 0, l-1; low <= hig; {
		mid := (low + hig) >> 1
		// fmt.Println(low, hig, mid)
		// fmt.Println(nums[low], nums[hig], nums[mid])
		if target == nums[mid] || target == nums[low] || target == nums[hig] {
			return true
		}

		if nums[mid] == nums[low] {
			low++
			continue
		}
		if nums[mid] == nums[hig] {
			hig--
			continue
		}

		// 左边有序
		if nums[mid] > nums[low] {
			// 在有序部分
			if nums[low] < target && target < nums[mid] {
				hig = mid - 1
			} else {
				low = mid + 1
			}

		} else {
			// 在有序部分
			if nums[hig] > target && target > nums[mid] {
				low = mid + 1
			} else {
				hig = mid - 1
			}
		}
	}
	return false
}

// Largest rectangle in histogram
func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}
	// 搜索数组,以当前高度,向左右寻找边界
	// 边界必定是小于等于当前高度所对应的高度对应的index
	// 左右边界的差值(边界肯定不包含在内)需要减一才是宽度

	left, right, stack := make([]int, l), make([]int, l), []int{}
	// 寻找左边界
	for i := 0; i < l; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		// 当前位置入栈
		stack = append(stack, i)
	}
	stack = []int{}
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = l
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < l; i++ {
		ans = max(ans, ((right[i] - left[i] - 1) * heights[i]))
	}
	return ans
}

// Maximal rectangle
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	maxArea := math.MinInt
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 通过寻找最大矩形的题解来解决当前问题
	largestRectangleArea := func(heights []int) int {
		l := len(heights)
		if l == 0 {
			return 0
		}
		if l == 1 {
			return heights[0]
		}
		// 搜索数组,以当前高度,向左右寻找边界
		// 边界必定是小于等于当前高度所对应的高度对应的index
		// 左右边界的差值(边界肯定不包含在内)需要减一才是宽度

		left, right, stack := make([]int, l), make([]int, l), []int{}
		// 寻找左边界
		for i := 0; i < l; i++ {
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				left[i] = -1
			} else {
				left[i] = stack[len(stack)-1]
			}
			// 当前位置入栈
			stack = append(stack, i)
		}
		stack = []int{}
		for i := l - 1; i >= 0; i-- {
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				right[i] = l
			} else {
				right[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		ans := 0
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		for i := 0; i < l; i++ {
			ans = max(ans, ((right[i] - left[i] - 1) * heights[i]))
		}
		return ans
	}

	for row := 0; row < n; row++ {
		input := make([]int, m)
		for col := 0; col < m; col++ {
			// 构建求解入参
			if row == 0 {
				if matrix[row][col] == '1' {
					input[col] = 1
				} else {
					input[col] = 0
				}
			} else {
				if matrix[row][col] == '0' {
					input[col] = 0
				} else {
					for now := row; now >= 0; now-- {
						if matrix[now][col] == '0' {
							break
						}
						input[col]++
					}
				}
			}
		}
		maxArea = max(maxArea, largestRectangleArea(input))
	}
	return maxArea
}
