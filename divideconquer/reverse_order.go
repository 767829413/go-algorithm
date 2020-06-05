package divideconquer

var num int

func Count(arr []int, n int) int {
	num = 0
	mergeSortCounting(arr, 0, n-1)
	return num
}

func mergeSortCounting(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := start + ((end - start) >> 1)
	mergeSortCounting(arr, start, mid)
	mergeSortCounting(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(a []int, start int, mid int, end int) {
	tmp := make([]int, end-start+1)
	index := 0
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if a[i] <= a[j] {
			tmp[index] = a[i]
			i++
		} else {
			num += mid - i + 1
			tmp[index] = a[j]
			j++
		}
		index++
	}
	for i <= mid {
		tmp[index] = a[i]
		i++
		index++
	}
	for j <= end {
		tmp[index] = a[j]
		j++
		index++
	}
	copy(a[start:end+1], tmp)
}
