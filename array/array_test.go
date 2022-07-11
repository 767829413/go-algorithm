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
	expected := []int{1, 2}
	assert.Equal(expected, twoSum(nums, target))
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)
	nums1 := []int{2, 3, 5}
	nums2 := []int{1, 4, 7, 9}
	expected := 4.0
	assert.Equal(expected, findMedianSortedArrays(nums1, nums2))
}

func TestMaxArea(t *testing.T) {
	assert := assert.New(t)
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	assert.Equal(expected, maxArea(height))
}

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	assert.Equal(expected, threeSum(nums))
}

func TestThreeSumClosest(t *testing.T) {
	assert := assert.New(t)
	// nums := []int{-1, 2, 1, -4}
	// approachNumber := 1
	// expected := 2
	nums := []int{-1, 2, 1, -4}
	approachNumber := 1
	expected := 2
	assert.Equal(expected, threeSumClosest(nums, approachNumber))
}

func TestFourSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 0, -1, 0, -2, 2}
	approachNumber := 0
	expected := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	// nums := []int{2, 2, 2, 2}
	// approachNumber := 8
	// expected := [][]int{}
	assert.Equal(expected, fourSum(nums, approachNumber))
}

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	targetNums := []int{0, 1, 2, 3, 4}
	assert.Equal(expected, removeDuplicates(nums))
	for k, v := range targetNums {
		assert.Equal(v, nums[k])
	}
}
func TestRemoveElement(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeNum := 2
	expected := 5
	targetNums := []int{0, 1, 4, 0, 3, 0, 4, 2}
	// nums := []int{2}
	// removeNum := 2
	// expected := 0
	// targetNums := []int{2}
	assert.Equal(expected, removeElement(nums, removeNum))
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
	expected := 4
	assert.Equal(expected, search(nums, target))
}

func TestSearchRange(t *testing.T) {
	assert := assert.New(t)
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	assert.Equal(expected, searchRange(nums, target))
}

func TestSearchInsert(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	assert.Equal(expected, searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 2
	expected = 1
	assert.Equal(expected, searchInsert(nums, target))
}

func TestIsValidSudoku(t *testing.T) {
	assert := assert.New(t)
	nums := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	expected := true
	assert.Equal(expected, isValidSudoku(nums))
	nums = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	expected = false
	assert.Equal(expected, isValidSudoku(nums))
}

func TestSolveSudoku(t *testing.T) {
	assert := assert.New(t)
	nums := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	expected := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
	solveSudoku(nums)
	assert.Equal(expected, nums)
}

func TestCombinationSum(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 3, 5}
	target := 8
	expected := [][]int{
		{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
	}
	assert.Equal(expected, combinationSum(nums, target))
}

func TestCombinationSum2(t *testing.T) {
	assert := assert.New(t)
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	expected := [][]int{
		{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
	}
	assert.Equal(expected, combinationSum2(nums, target))
}

func TestFirstMissingPositive(t *testing.T) {
	assert := assert.New(t)
	nums := []int{3, 4, -1, 1}
	expected := 2
	assert.Equal(expected, firstMissingPositive(nums))
}

func TestTrap(t *testing.T) {
	assert := assert.New(t)
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	assert.Equal(expected, trap(nums))
}

func TestJump(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 3, 1, 1, 4}
	expected := 2
	assert.Equal(expected, jump(nums))
}

func TestPermute(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	assert.Equal(expected, permute(nums))
}

func TestPermuteUnique(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 1, 2}
	expected := [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}
	assert.Equal(expected, permuteUnique(nums))
}

func TestRotate(t *testing.T) {
	assert := assert.New(t)
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	rotate(nums)
	assert.Equal(expected, nums)
}

func TestGroupAnagrams(t *testing.T) {
	assert := assert.New(t)
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"tan", "nat"},
		{"eat", "tea", "ate"},
	}
	for _, v := range groupAnagrams(strs) {
		assert.Contains(expected, v)
	}
}

func TestSolveNQueens(t *testing.T) {
	assert := assert.New(t)
	input := 4
	expected := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	for _, v := range solveNQueens(input) {
		assert.Contains(expected, v)
	}
}

func TestMaxSubArray(t *testing.T) {
	assert := assert.New(t)
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	assert.Equal(expected, maxSubArray(input))

}

func TestSpiralOrder(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	assert.Equal(expected, spiralOrder(input))
}

func TestCanJump(t *testing.T) {
	assert := assert.New(t)
	input := []int{1, 1, 1, 0}
	expected := true
	assert.Equal(expected, canJump(input))
}

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{
		{1, 10},
		{2, 5},
		{3, 6},
		{4, 7},
	}
	expected := [][]int{
		{1, 10},
	}
	assert.Equal(expected, merge(input))
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{4, 8}
	expected := [][]int{
		{1, 2},
		{3, 10},
		{12, 16},
	}
	assert.Equal(expected, insert(input, newInterval))
}

func TestGenerateMatrix(t *testing.T) {
	assert := assert.New(t)
	input := 3
	expected := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	assert.Equal(expected, generateMatrix(input))
}

func TestUniquePathsWithObstacles(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	expected := 2
	assert.Equal(expected, uniquePathsWithObstacles(input))
}

func TestMinPathSum(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	expected := 7
	assert.Equal(expected, minPathSum(input))
}
