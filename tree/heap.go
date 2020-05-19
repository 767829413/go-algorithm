package tree

type Heap struct {
	v     []int //数组
	n     int   //容量
	count int   //已存数据数量
}

func NewHeap(n int) *Heap {
	return &Heap{
		v: make([]int, n+1),
		n: n,
	}
}

//自下向上交换
func (h *Heap) Insert(d int) {
	if h.count >= h.n {
		return
	}
	h.count++
	h.v[h.count] = d
	i := h.count
	for (i>>1) > 0 && h.v[i] > h.v[i>>1] {
		h.v[i], h.v[i>>1] = h.v[i>>1], h.v[i]
		i = i >> 1
	}
}

func (h *Heap) DelMax() {
	if h.count == 0 {
		return
	}
	h.v[1], h.v[h.count] = h.v[h.count], 0
	h.count--
	Heapify(h.v, 1, h.count)
}

/**
构建堆
在利用数组存储时(根节点0不用),从下标1开始到n/2,都是非叶子节点,n/2到n均为叶子节点
堆是完全二叉树，求最后的非叶子节点即是求最大的叶子节点的父节点
最大的叶子节点下标为n，他的父节点为n/2，这是最后一个非叶子节点，所以n/2+1到n都是叶子节点
*/
func BuildHeap(arr []int, n int) {
	for i := n >> 1; i > 1; i-- {
		Heapify(arr, i, n)
	}
}

func HeapSort(arr []int, n int) {
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

func IsBigTopHeap(arr []int, n int) bool {
	i := 1
	for i <= n {
		left := i << 1
		if left <= n && arr[i] < arr[left] {
			return false
		}
		right := i<<1 + 1
		if right <= n && arr[i] < arr[right] {
			return false
		}
		i++
	}
	return true
}
