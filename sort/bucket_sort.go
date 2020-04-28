package sort

func getMax(a []int, num int) int {
	max := a[0]
	for i := 1; i < num; i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func BucketSort(a []int) {
	count := len(a)
	if count <= 1 {
		return
	}
	max := getMax(a, count)
	buckets := make([][]int, count)
	for i := 0; i < count; i++ {
		index := (a[i] * (count - 1)) / max
		buckets[index] = append(buckets[index], a[i])
	}
	pos := 0
	for i := 0; i < count; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			Quicksort(buckets[i])
			copy(a[pos:], buckets[i])
			pos += bucketLen
		}
	}
}

func BucketSortSimple(a []int) {
	count := len(a)
	if count <= 1 {
		return
	}
	max := getMax(a, count)
	bucketsLen := max + 1
	buckets := make([]int, bucketsLen)
	for i := 0; i < count; i++ {
		buckets[a[i]]++
	}
	tmp := make([]int, 0)
	for i := 0; i < bucketsLen; i++ {
		for buckets[i] > 0 {
			tmp = append(tmp, i)
			buckets[i]--
		}
	}
	copy(a, tmp)
}
