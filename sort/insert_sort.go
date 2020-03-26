package sort

func InsertSort(arr []int) {
	count := len(arr)
	if count <= 1 {
		return
	}
	for i := 1; i < count; i++ {
		val := arr[i]
		index := i - 1
		for ; index >= 0; index-- {
			if arr[index] > val {
				arr[index+1] = arr[index] //在有序的序列对比查找位置
			} else {
				break
			}
		}
		arr[index+1] = val //找到位置后插入(跳出循环是index已经--操作了)
	}
}
