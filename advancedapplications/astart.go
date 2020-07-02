package advancedapplications

import (
	"container/list"
	"fmt"
	"math"
)

type SpaGraphAStart struct {
	adj      []*list.List //邻接表
	v        int          //顶点个数
	vertexes []*VertexAStart
}

func NewSpaGraphAStart(v int) *SpaGraphAStart {
	adj := make([]*list.List, v)
	vertexes := make([]*VertexAStart, v)
	for i := 0; i < v; i++ {
		adj[i] = list.New()
	}
	return &SpaGraphAStart{
		adj:      adj,
		v:        v,
		vertexes: vertexes,
	}
}
func (s *SpaGraphAStart) AddVetex(i, x, y int) {
	s.vertexes[i] = NewVertexAStart(i, x, y)
}
func (s *SpaGraphAStart) AddEdge(start, end int, wight int) {
	s.adj[start].PushBack(NewSpaEdgeAStart(start, end, wight))
}

func (s *SpaGraphAStart) hManhattan(v1, v2 *VertexAStart) int {
	return int(math.Abs(float64(v1.x-v2.x)) + math.Abs(float64(v1.y-v2.y)))
}

// 计算 从顶点start到顶点end的路径
func (s *SpaGraphAStart) AStar(start, end int) {
	//pressPath用来还原最短路径
	pressPath := make([]int, s.v)
	//优先级队列(小顶堆)
	queue := NewPriorityQueueAStart(s.v)
	//标记是否进入过队列
	inQueue := make([]bool, s.v)
	s.vertexes[start].dist = 0
	s.vertexes[start].f = 0
	queue.add(s.vertexes[start])
	inQueue[start] = true
	for !queue.isEmpty() {
		//取堆顶元素并删除
		minVertex := queue.poll()
		//取出minVetex相连的边
		for edge := s.adj[minVertex.id].Front(); edge != nil; edge = edge.Next() {
			//minVertex-->nextVertex
			nextVertex := s.vertexes[edge.Value.(*SpaEdgeAStart).eid]
			curDist := minVertex.dist + edge.Value.(*SpaEdgeAStart).w
			if curDist < nextVertex.dist {
				//更新next的dis
				nextVertex.dist = curDist
				nextVertex.f = curDist + s.hManhattan(nextVertex, s.vertexes[end])
				pressPath[nextVertex.id] = minVertex.id
				if inQueue[nextVertex.id] {
					//更新队列中的dist值
					queue.update(nextVertex)
				} else {
					queue.add(nextVertex)
					inQueue[nextVertex.id] = true
				}
			}
			//最短路径产生了
			if minVertex.id == end {
				queue.clear()
				break
			}
		}
	}
	//输出最短路径
	fmt.Print(start)
	s.print(start, end, pressPath)
	fmt.Println()
	fmt.Println(pressPath)
}

func (s *SpaGraphAStart) print(start, end int, pressPath []int) {
	if start == end {
		return
	}
	s.print(start, pressPath[end], pressPath)
	fmt.Print("->", end)
}

type SpaEdgeAStart struct {
	sid int //起始点id
	eid int //终结点id
	w   int //权重
}

func NewSpaEdgeAStart(sid, eid, w int) *SpaEdgeAStart {
	return &SpaEdgeAStart{
		sid: sid,
		eid: eid,
		w:   w,
	}
}

type VertexAStart struct {
	id   int //顶点id
	dist int //起始点到改点的距离
	f    int //估值
	x    int //横坐标
	y    int //纵坐标
}

func NewVertexAStart(id, x, y int) *VertexAStart {
	return &VertexAStart{
		id: id,
		x:  x,
		y:  y,
	}
}

type PriorityQueueAStart struct {
	nodes []*VertexAStart
	count int
	index map[int]int
}

func NewPriorityQueueAStart(v int) *PriorityQueueAStart {
	return &PriorityQueueAStart{
		nodes: make([]*VertexAStart, v+1),
		index: make(map[int]int, v),
	}
}

func (p *PriorityQueueAStart) poll() *VertexAStart {
	v := p.nodes[1]
	p.nodes[1], p.nodes[p.count] = p.nodes[p.count], nil
	p.count--
	p.heapifyByUpToDown()
	return v
}

func (p *PriorityQueueAStart) add(v *VertexAStart) {
	p.count++
	p.nodes[p.count] = v
	p.heapifyByDownToUp()
}
func (p *PriorityQueueAStart) update(v *VertexAStart) {
	//找到更新的节点
	p.nodes[p.index[v.id]] = v
	//自下往上堆化
	p.heapifyByDownToUp()
}

func (p *PriorityQueueAStart) heapifyByDownToUp() {
	i := p.count
	for i>>1 > 0 && p.nodes[i].f < p.nodes[i>>1].f {
		p.index[p.nodes[i].id] = p.nodes[i>>1].id
		p.index[p.nodes[i>>1].id] = p.nodes[i].id
		p.nodes[i], p.nodes[i>>1] = p.nodes[i>>1], p.nodes[i]
		i = i >> 1
	}
}

func (p *PriorityQueueAStart) heapifyByUpToDown() {
	start := 1
	for {
		minIndex := start
		left := start << 1
		right := left + 1
		leftFlag := false
		rightFlag := false
		if left <= p.count && p.nodes[left].f < p.nodes[start].f {
			leftFlag = true
			minIndex = left
		}
		if right <= p.count && p.nodes[right].f < p.nodes[start].f {
			rightFlag = true
			minIndex = right
		}
		if minIndex == start {
			break
		}
		if leftFlag && rightFlag {
			minIndex = right
			if p.nodes[left].f < p.nodes[right].f {
				minIndex = left
			}
		}
		p.index[p.nodes[start].id] = p.nodes[minIndex].id
		p.index[p.nodes[minIndex].id] = p.nodes[start].id
		p.nodes[start], p.nodes[minIndex] = p.nodes[minIndex], p.nodes[start]
		start = minIndex
	}
}
func (p *PriorityQueueAStart) isEmpty() bool {
	return p.count == 0
}

func (p *PriorityQueueAStart) clear() {
	p.nodes = make([]*VertexAStart, len(p.nodes))
	p.count = 0
}
