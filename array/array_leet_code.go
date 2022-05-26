package array

import (
	"math"
	"sort"
)

// 两数之和
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

// 寻找两个正序数组的中位数
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

// 盛最多水的容器
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

// 三数之和
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
