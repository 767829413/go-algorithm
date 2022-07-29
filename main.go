package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	// 定义解集合
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	// 假设 n = 4
	// 当 i = 1 时候 左边没有,右边3个均大于 1 dp[1] += dp[3]
	// 当 i = 2 时候 左边只有 1,右边2个均大于 2 dp[2] += dp[1] * dp[2]
	// 当 i = 3 时候 左边只有 1 2 ,右边1个均大于 3 dp[3] += dp[2] * dp[1]

	// i 左边 i-1 个比 i 小,右边 n-i 比 i 大 dp[i] += dp[i-1] * dp[n-i]
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}
