package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs([]int{}, nums, 0)
	return res
}
func dfs(temp, num []int, start int) {
	fmt.Println(start, temp)
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(num); i++ {
		if i > start && num[i] == num[i-1] {
			continue
		}
		temp = append(temp, num[i])
		dfs(temp, num, i+1)
		temp = temp[:len(temp)-1]
	}
}
