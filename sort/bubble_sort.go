package sort

func BubbleSort(arr []int) {
	count := len(arr)
	if count <= 1 {
		return
	}
	for i := 0; i < count; i++ {
		flag := false
		for j := 0; j < count-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
