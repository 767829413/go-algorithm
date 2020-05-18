package tree

import "fmt"

func FullRange(arr []int, n, k int, num *int) {
	if k == 1 {
		*num++
		fmt.Println(arr)
		return
	}
	for i := 0; i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]
		FullRange(arr, n, k-1, num)
		arr[i], arr[k-1] = arr[k-1], arr[i]
	}
}

/**
1 个细胞的生命周期是 3 小时，1 小时分裂一次。求 n 小时后，容器内有多少细胞？
 */
func CellBreed(n int) int {
	if n < 1 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 4
	}
	if n == 3 {
		return 7
	}
	return 2*CellBreed(n-1) - CellBreed(n-4)
}
