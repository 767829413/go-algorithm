package array

import (
	"math"
)

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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := l1 + l2
	lmax1, lmax2, rmin1, rmin2, c1, c2, lo, hi := 0, 0, 0, 0, 0, 0, 0, l1<<1
	for lo < hi {
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
