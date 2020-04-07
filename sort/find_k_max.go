package sort

/**
分区的思想
第k大就是表示前面有k-1个数比它大
 */
func FindKMax(arr []int, k int) int {
	var find int
	if len(arr) < k || k < 0 {
		return 0
	}
	left := 0
	right := len(arr) - 1
	findK(arr, left, right, k, &find)
	if find == 0 {
		return arr[k-1]
	}
	return find
}

func findK(arr []int, left int, right int, k int, find *int) {
	if left >= right {
		return
	}
	//分区
	pt := partitionFind(arr, left, right, k, find)
	findK(arr, left, pt-1, k, find)
	findK(arr, pt+1, right, k, find)
}

func partitionFind(arr []int, left, right int, k int, find *int) int {
	pivot := arr[right]
	i := left
	j := left
	for ; j < right; j = j + 1 {
		if arr[j] > pivot {
			swap(arr, i, j);
			i++
		}
	}
	if i+1 == k {
		*find = pivot
	}
	swap(arr, i, right);
	return i
}
