package sort

func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left := 0
	right := len(arr) - 1
	divideQS(arr, left, right)
}

func divideQS(arr []int, left, right int) {
	if left >= right {
		return
	}
	//分区
	pt := partition(arr, left, right)
	//递归
	divideQS(arr, left, pt-1)
	divideQS(arr, pt+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j = j + 1 {
		if arr[j] < pivot {
			swap(arr, i, j);
			i++
		}
	}
	swap(arr, i, right);
	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
