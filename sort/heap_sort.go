package sort

func BuildHeap(arr []int, n int) {
	for i := n >> 1; i > 1; i-- {
		Heapify(arr, i, n)
	}
}

func HeapSort(arr []int) {
	n := len(arr) - 1
	BuildHeap(arr, n)
	i := n
	for i > 1 {
		arr[i], arr[1] = arr[1], arr[i]
		i--
		Heapify(arr, 1, i)
	}
}

//堆化
func Heapify(arr []int, start, count int) {
	for {
		maxIndex := start
		left := start << 1
		right := (start << 1) + 1
		flagLeft := false
		flagRight := false
		if left <= count && arr[start] < arr[left] {
			flagLeft = true
			maxIndex = left
		}
		if right <= count && arr[start] < arr[right] {
			flagRight = true
			maxIndex = right
		}
		if maxIndex == start {
			break
		}
		if flagLeft && flagRight {
			if arr[left] > arr[right] {
				maxIndex = left
			}
		}
		arr[maxIndex], arr[start] = arr[start], arr[maxIndex]
		start = maxIndex
	}
}
