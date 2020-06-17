package dynamicprogramming

/**
我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？
比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，
它的最长递增子序列就是 2, 3, 5, 7，
所以最长递增子序列的长度是 4。
 */

type LISL struct {
	arr []int
	n   int
	max int
}

func NewLISL(arr []int, n int) *LISL {
	return &LISL{
		arr: arr,
		n:   n,
		max: ^int(^uint(0) >> 1),
	}
}

//使用动态规划解决
func (l *LISL) GetMaxLength() int {
	return l.toGet()
}

func (l *LISL) toGet() int {
	if l.n == 0 {
		return 0
	}
	dp, res := make([]int, l.n), 0
	for i := 0; i < l.n; i++ {
		dp[i] = 1
	}
	for i := 0; i < l.n; i++ {
		for j := 0; j < i; j++ {
			if l.arr[i] > l.arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
