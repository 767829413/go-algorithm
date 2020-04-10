package sort

func SelectSort(arr []int) {
	var min int
	count := len(arr)
	if count <= 1 {
		return
	}
	for i := 0; i < count; i++ {
		min = i
		for j := i + 1; j < count; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
