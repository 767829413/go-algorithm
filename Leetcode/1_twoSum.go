package Leetcode

func twoSum(nums []int, target int) []int {
	result := make([]int, 2, 2)
	n := len(nums)
	c := make(map[int]int, n)
	for i := 0; i < n; i++ {
		cur := target - nums[i]
		if index, ok := c[cur]; ok {
			result[0] = index
			result[1] = i
		}
		c[nums[i]] = i
	}
	return result
}
