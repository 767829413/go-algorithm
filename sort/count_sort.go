package sort

func CountSort(arr []int) {
	count := len(arr)
	if count <= 1 {
		return
	}
	max := arr[0]
	for i := 1; i < count; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	bucket := make([]int, max+1, max+1)
	for i := 0; i < count; i++ {
		bucket[arr[i]]++
	}

	for i := 1; i <= max; i++ {
		bucket[i] += bucket[i-1]
	}

	tmp := make([]int, count, count)
	for i := range arr {
		index := bucket[arr[i]] - 1
		tmp[index] = arr[i]
		bucket[arr[i]]--
	}
	copy(arr, tmp)
}
