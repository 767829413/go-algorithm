package tree

type pnode struct {
	v     interface{}
	level int
}

func NewPnode(v interface{}, level int) *pnode {
	return &pnode{
		v:     v,
		level: level,
	}
}

type PQueue struct {
	heap  []*pnode
	n     int
	count int
}

func NewPQueue(n int) *PQueue {
	return &PQueue{
		heap: make([]*pnode, n+1, n+1),
		n:    n,
	}
}

//入队
func (p *PQueue) Push(node *pnode) bool {
	if p.count > p.n {
		return false
	}
	p.count++
	p.heap[p.count] = node
	start := p.count
	for start>>1 > 0 && p.heap[start].level > p.heap[start>>1].level {
		p.heap[start], p.heap[start>>1] = p.heap[start>>1], p.heap[start]
		start = start >> 1
	}
	return true
}

//出队
func (p *PQueue) Pop() (*pnode, bool) {
	if p.count == 0 {
		return nil, false
	}
	v := p.heap[1]
	p.heap[1], p.heap[p.count] = p.heap[p.count], p.heap[1]
	p.count--
	p.heapify(1)
	return v, true
}

//堆化
func (p *PQueue) heapify(start int) {
	for {
		maxIndex := start
		left := maxIndex << 1
		right := maxIndex<<1 + 1
		flagLeft := false
		flagRight := false
		if left <= p.count && p.heap[maxIndex].level < p.heap[left].level {
			flagLeft = true
			maxIndex = left
		}
		if right <= p.count && p.heap[maxIndex].level < p.heap[right].level {
			flagRight = true
			maxIndex = right
		}
		if flagLeft && flagRight && p.heap[left].level > p.heap[right].level {
			maxIndex = left
		}
		if maxIndex == start {
			break
		}
		p.heap[maxIndex], p.heap[start] = p.heap[start], p.heap[maxIndex]
		start = maxIndex
	}
}

func (p *PQueue) IsBigTopHeap() bool {
	i := 1
	for i <= p.count {
		left := i << 1
		if left <= p.count && p.heap[i].level < p.heap[left].level {
			return false
		}
		right := i<<1 + 1
		if right <= p.count && p.heap[i].level < p.heap[right].level {
			return false
		}
		i++
	}
	return true
}
