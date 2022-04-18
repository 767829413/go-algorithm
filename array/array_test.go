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
	target := 9
	assert.Equal([]int{0, 1}, twoSum(nums, target))
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	target := 2.5
	assert.Equal(target, findMedianSortedArrays(nums1, nums2))
}
