package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInsert(t *testing.T) {
	capacity := 10
	var err error
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-3; i++ {
		err = arr.Insert(uint(i), i+2)
		if err != nil {
			t.Fatal(err)
		}
	}
	arr.Print()
	err = arr.Insert(uint(capacity+1), 10)
	arr.Print()
	t.Log(err)
}

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 7, 11, 15}
	target := 18
	assert.Equal([]int{1, 2}, twoSum(nums, target))
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)
	nums1 := []int{2, 3, 5}
	nums2 := []int{1, 4, 7, 9}
	target := 4.0
	assert.Equal(target, findMedianSortedArrays(nums1, nums2))
}

func TestMaxArea(t *testing.T) {
	assert := assert.New(t)
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 49
	assert.Equal(target, maxArea(height))
}

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	assert.Equal(target, threeSum(nums))
}

func TestThreeSumClosest(t *testing.T) {
	assert := assert.New(t)
	// nums := []int{-1, 2, 1, -4}
	// approachNumber := 1
	// target := 2
	nums := []int{-1, 2, 1, -4}
	approachNumber := 1
	target := 2
	assert.Equal(target, threeSumClosest(nums, approachNumber))
}

func TestFourSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 0, -1, 0, -2, 2}
	approachNumber := 0
	target := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	// nums := []int{2, 2, 2, 2}
	// approachNumber := 8
	// target := [][]int{}
	assert.Equal(target, fourSum(nums, approachNumber))
}

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	target := 5
	targetNums := []int{0, 1, 2, 3, 4}
	assert.Equal(target, removeDuplicates(nums))
	for k, v := range targetNums {
		assert.Equal(v, nums[k])
	}
}
func TestRemoveElement(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeNum := 2
	target := 5
	targetNums := []int{0, 1, 4, 0, 3, 0, 4, 2}
	// nums := []int{2}
	// removeNum := 2
	// target := 0
	// targetNums := []int{2}
	assert.Equal(target, removeElement(nums, removeNum))
	assert.Equal(targetNums, nums)
}

func TestNextPermutation(t *testing.T) {
	assert := assert.New(t)
	nums := []int{3, 2, 1}
	targetNums := []int{1, 2, 3}
	nextPermutation(nums)
	assert.Equal(targetNums, nums)
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	output := 4
	assert.Equal(output, search(nums, target))
}

func TestSearchRange(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	output := []int{3, 4}
	assert.Equal(output, searchRange(nums, target))
}

func TestSearchInsert(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 3, 5, 6}
	target := 5
	output := 2
	assert.Equal(output, searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 2
	output = 1
	assert.Equal(output, searchInsert(nums, target))
}
