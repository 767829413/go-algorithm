package greedy

import "fmt"

type HuffManTree struct {
	Root *node
}

func NewHuffManTree(str []rune) *HuffManTree {
	h := &HuffManTree{}
	build(str, h)
	return h
}

func (h *HuffManTree) Print() {
	cur := h.Root
	path := "0"
	p(cur, path)
}

func p(node *node, path string) {
	if node == nil {
		return
	}
	if node.value != '/' {
		fmt.Println(string(node.value), path)
	}
	p(node.left, path+"0")
	p(node.right, path+"1")
}

//构建
func build(str []rune, h *HuffManTree) {
	num := len(str)
	//统计字符频率
	res := countCharNum(str, num)
	//构建最小堆(优先级队列)
	heap := newHeap(len(res))
	for k, v := range res {
		heap.insert(newNode(k, v, true))
	}
	for heap.getCount() > 1 {
		n1 := heap.popTop()
		n2 := heap.popTop()
		node := newNode('/', n1.freq+n2.freq, false)
		n1.wight = 0
		n2.wight = 1
		node.left = n1
		node.right = n2
		heap.insert(node)
	}
	nodeRes := heap.popTop()
	if num == 1 {
		h.Root = newNode('/', nodeRes.freq+1, false)
		h.Root.right = nodeRes
	} else {
		h.Root = nodeRes
	}
}

type node struct {
	freq        int
	wight       int
	value       rune
	left, right *node
	isEnd       bool
}

func newNode(v rune, freq int, isEnd bool) *node {
	return &node{
		value: v,
		freq:  freq,
		isEnd: isEnd,
	}
}

type heap struct {
	v     []*node //数组
	n     int     //容量
	count int     //已存数据数量
}

func newHeap(n int) *heap {
	return &heap{
		v: make([]*node, n+1),
		n: n,
	}
}

func (h *heap) isEmpty() bool {
	return h.count == 0
}

func (h *heap) getCount() int {
	return h.count
}

func (h *heap) insert(node *node) {
	if h.count >= h.n {
		return
	}
	h.count++
	h.v[h.count] = node
	i := h.count
	for (i>>1) > 0 && h.v[i].freq < h.v[i>>1].freq {
		h.v[i], h.v[i>>1] = h.v[i>>1], h.v[i]
		i = i >> 1
	}
}

func (h *heap) popTop() *node {
	v := h.v[1]
	h.v[1], h.v[h.count] = h.v[h.count], nil
	h.count--
	heapify(h.v, 1, h.count)
	return v
}

func (h *heap) print() {
	for i := 1; i <= h.count; i++ {
		fmt.Println(string(h.v[i].value), h.v[i].freq)
	}
}

func heapify(arr []*node, start, count int) {
	for {
		minIndex := start
		left := start << 1
		right := (start << 1) + 1
		flagLeft := false
		flagRight := false
		if left <= count && arr[start].freq > arr[left].freq {
			flagLeft = true
			minIndex = left
		}
		if right <= count && arr[start].freq > arr[right].freq {
			flagRight = true
			minIndex = right
		}
		if minIndex == start {
			break
		}
		if flagLeft && flagRight {
			if arr[left].freq < arr[right].freq {
				minIndex = left
			} else {
				minIndex = right
			}
		}
		arr[minIndex], arr[start] = arr[start], arr[minIndex]
		start = minIndex
	}
}

func countCharNum(str []rune, n int) (result map[rune]int) {
	result = make(map[rune]int, n)
	for i := 0; i < n; i++ {
		if v, ok := result[str[i]]; ok {
			result[str[i]] = v + 1
		} else {
			result[str[i]] = 1
		}
	}
	return result
}
