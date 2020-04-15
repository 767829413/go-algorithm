package search

func EasyLoopBinarySearch(arr []int, find int, index *int) {
	count := len(arr)
	if count < 1 {
		*index = -1
	}
	lo := 0
	hi := count - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		switch true {
		case arr[mid] == find:
			*index = mid
			return
		case arr[mid] < find:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	*index = -1
}

func EasyRecursiveBinarySearch(arr []int, find int, index *int) {
	count := len(arr)
	if count < 1 {
		*index = -1
	}
	lo := 0
	hi := count - 1
	*index = recursiveItem(arr, find, lo, hi)
}

func recursiveItem(arr []int, find, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + ((hi - lo) >> 1)
	switch true {
	case arr[mid] == find:
		return mid
	case arr[mid] < find:
		return recursiveItem(arr, find, mid+1, hi)
	default:
		return recursiveItem(arr, find, lo, mid-1)
	}
}

func FindFirstEqual(arr []int, find int, index *int) {
	count := len(arr)
	if count < 1 {
		*index = -1
	}
	lo := 0
	hi := count - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] < find {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if lo < count && arr[lo] == find {
		*index = lo
	}
}
//查找指定值
func FindEndEqual(arr []int, find int, index *int) {
	count := len(arr)
	if count < 1 {
		*index = -1
	}
	lo := 0
	hi := count - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] > find {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if hi < count && arr[hi] == find {
		*index = hi
	}
}

//查找第一个大于等于指定值
func FindFirstGe(arr []int, find int, index *int) {
	count := len(arr)
	lo := 0
	hi := count - 1
	if count < 1 || arr[hi] < find {
		*index = -1
		return
	}
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] >= find {
			if mid == 0 || arr[mid-1] < find {
				*index = mid
				break
			} else {
				hi = mid - 1
			}
		} else {
			lo = mid + 1
		}
	}
}

//查找最后一个大于等于指定值
func FindEndGe(arr []int, find int, index *int) {
	count := len(arr)
	lo := 0
	hi := count - 1
	mid := 0
	if count < 1 {
		*index = -1
		return
	}
	for lo <= hi {
		mid = lo + ((hi - lo) >> 1)
		if arr[mid] <= find {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if hi < count {
		if arr[hi] == find {
			*index = hi
		} else {
			*index = mid
		}
	}
}

//查找第一个小于等于指定值
func FindFirstLe(arr []int, find int, index *int) {
	count := len(arr)
	lo := 0
	hi := count - 1
	mid := 0
	if count < 1 {
		*index = -1
		return
	}
	for lo <= hi {
		mid = lo + ((hi - lo) >> 1)
		if find >= arr[mid] {
			*index = lo
			break
		} else {
			hi = mid - 1
		}
	}
}

//查找最后一个小于等于指定值
func FindEndLe(arr []int, find int, index *int) {
	count := len(arr)
	lo := 0
	hi := count - 1
	mid := 0
	if count < 1 {
		*index = -1
		return
	}
	for lo <= hi {
		mid = lo + ((hi - lo) >> 1)
		if find >= arr[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	*index = hi
}
