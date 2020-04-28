package sort

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left := 0
	right := len(arr) - 1
	divideMS(arr, left, right)
}

func divideMS(arr []int, left, right int) {
	if left >= right {
		return
	}
	//取left right 中间位置
	middle := (left + right) >> 1
	//分治递归
	divideMS(arr, left, middle)
	divideMS(arr, middle+1, right)
	//合并
	conquer(arr, left, middle, right)
}

func conquer(arr []int, left, middle, right int) {
	var tempArr []int
	pl := left
	pr := middle + 1
	for pl <= middle && pr <= right {
		if arr[pl] <= arr[pr] {
			tempArr = append(tempArr, arr[pl])
			pl++
		} else {
			tempArr = append(tempArr, arr[pr])
			pr++
		}
	}
	for pr <= right {
		tempArr = append(tempArr, arr[pr])
		pr++
	}
	for pl <= middle {
		tempArr = append(tempArr, arr[pl])
		pl++
	}
	//复制数据
	copy(arr[left:right+1], tempArr)
}
