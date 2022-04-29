package array

import (
	"github.com/767829413/go-algorithm/utilcom"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if index, ok := m[tmp]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, num1mid, num2mid := 0, l1, (l1+l2+1)>>1, 0, 0
	for low < high {
		num1mid = low + ((high - low) >> 1)
		num2mid = k - num1mid
		if num1mid > 0 && nums1[num1mid-1] > nums2[num2mid] {
			high = num1mid - 1
		} else if num1mid != l1 && nums2[num2mid-1] > nums1[num1mid] {
			low = num1mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if num1mid == 0 {
		midLeft = nums2[num2mid-1]
	} else if num2mid == 0 {
		midLeft = nums1[num1mid-1]
	} else {
		midLeft = utilcom.Max(nums1[num1mid-1], nums2[num2mid-1])
	}
	if (l1+l2)&1 == 1 {
		return float64(midLeft)
	}
	if num1mid == l1 {
		midRight = nums2[num2mid]
	} else if num2mid == l2 {
		midRight = nums1[num1mid]
	} else {
		midRight = utilcom.Min(nums1[num1mid], nums2[num2mid])
	}
	return float64(midLeft+midRight) / 2
}
