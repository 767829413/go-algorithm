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

func approachNumber(a, b, c int) int {
	l := math.Abs(float64(a) - float64(c))
	r := math.Abs(float64(b) - float64(c))
	if l > r {
		return b
	}
	return a
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
