package sort

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left := 0
	right := len(arr) - 1
	divide(arr, left, right)
}

func divide(arr []int, left, right int) {
	if left >= right {
		return
	}
	//取left right 中间位置
	middle := (left + right) / 2
	//分治递归
	divide(arr, left, middle)
	divide(arr, middle+1, right)
	//合并
	conquer(arr, left, middle, right)
}

func conquer(arr []int, left, middle, right int) {
	tempArr := []int{}
	pl := left
	pr := middle + 1
	for {
		if arr[pl] <= arr[pr] {
			tempArr = append(tempArr, arr[pl])
			pl++
		} else {
			tempArr = append(tempArr, arr[pr])
			pr++
		}
		if pl > middle {
			for pr <= right {
				tempArr = append(tempArr, arr[pr])
				pr++
			}
			break
		}
		if pr > right {
			for pl <= middle {
				tempArr = append(tempArr, arr[pl])
				pl++
			}
			break
		}
	}
	//复制数据
	for i := 0; left <= right; i, left = i+1, left+1 {
		arr[left] = tempArr[i]
	}
}
