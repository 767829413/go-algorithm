package advancedapplications

import (
	"container/list"
	"fmt"
)

/**
Shortest Path Algorithm 最短路径算法
1. 构建数据结构(有向有权图)
2. 构建有向有权图边的数据结构
3. 构建一个起点到顶点的距离
4. 构建一个优先级队列(小顶堆)
5. 使用dijkstra
 */
var intMAX = int(^uint(0) >> 1)

type SpaGraph struct {
	adj []*list.List //邻接表
	v   int          //顶点个数
}

func NewSpaGraph(v int) *SpaGraph {
	adj := make([]*list.List, v)
	for i := 0; i < v; i++ {
		adj[i] = list.New()
	}
	return &SpaGraph{
		adj: adj,
		v:   v,
	}
}

func (s *SpaGraph) AddEdge(start, end int, wight int) {
	s.adj[start].PushBack(NewSpaEdge(start, end, wight))
}

// 计算 start 到 end 的最短距离
func (s *SpaGraph) Dijkstra(start, end int) {
	//pressPath用来还原最短路径
	pressPath := make([]int, s.v)
	vertexes := make([]*Vertex, s.v)
	for i := 0; i < s.v; i++ {
		vertexes[i] = NewVertex(i, intMAX)
	}
	//优先级队列(小顶堆)
	queue := NewPriorityQueue(s.v)
	//标记是否进入过队列
	inQueue := make([]bool, s.v)
	vertexes[start].dist = 0
	queue.add(vertexes[start])
	inQueue[start] = true
	for !queue.isEmpty() {
		//取堆顶元素并删除
		minVertex := queue.poll()
		//最短路径产生了
		if minVertex.id == end {
			break
		}
		//取出minVetex相连的边
		for edge := s.adj[minVertex.id].Front(); edge != nil; edge = edge.Next() {
			//minVertex-->nextVertex
			nextVertex := vertexes[edge.Value.(*SpaEdge).eid]
			curDist := minVertex.dist + edge.Value.(*SpaEdge).w
			if curDist < nextVertex.dist {
				//更新next的dis
				nextVertex.dist = curDist
				pressPath[nextVertex.id] = minVertex.id
				if inQueue[nextVertex.id] {
					//更新队列中的dist值
					queue.update(nextVertex)
				} else {
					queue.add(nextVertex)
					inQueue[nextVertex.id] = true
				}
			}
		}
	}
	//输出最短路径
	fmt.Print(start)
	s.print(start, end, pressPath)
	fmt.Println()
	fmt.Println(pressPath)
}

func (s *SpaGraph) print(start, end int, pressPath []int) {
	if start == end {
		return
	}
	s.print(start, pressPath[end], pressPath)
	fmt.Print("->", end)
}

type SpaEdge struct {
	sid int //起始点id
	eid int //终结点id
	w   int //权重
}

func NewSpaEdge(sid, eid, w int) *SpaEdge {
	return &SpaEdge{
		sid: sid,
		eid: eid,
		w:   w,
	}
}

type Vertex struct {
	id   int //顶点id
	dist int //起始点到改点的距离
}

func NewVertex(id, dist int) *Vertex {
	return &Vertex{
		id:   id,
		dist: dist,
	}
}

type PriorityQueue struct {
	nodes []*Vertex
	count int
	index map[int]int
}

func NewPriorityQueue(v int) *PriorityQueue {
	return &PriorityQueue{
		nodes: make([]*Vertex, v+1),
		index: make(map[int]int, v),
	}
}

func (p *PriorityQueue) poll() *Vertex {
	v := p.nodes[1]
	p.nodes[1], p.nodes[p.count] = p.nodes[p.count], nil
	p.count--
	p.heapifyByUpToDown()
	return v
}

func (p *PriorityQueue) add(v *Vertex) {
	p.count++
	p.nodes[p.count] = v
	p.heapifyByDownToUp()
}
func (p *PriorityQueue) update(v *Vertex) {
	//找到更新的节点
	p.nodes[p.index[v.id]] = v
	//自下往上堆化
	p.heapifyByDownToUp()
}

func (p *PriorityQueue) heapifyByDownToUp() {
	i := p.count
	for i>>1 > 0 && p.nodes[i].dist < p.nodes[i>>1].dist {
		p.index[p.nodes[i].id] = p.nodes[i>>1].id
		p.index[p.nodes[i>>1].id] = p.nodes[i].id
		p.nodes[i], p.nodes[i>>1] = p.nodes[i>>1], p.nodes[i]
		i = i >> 1
	}
}

func (p *PriorityQueue) heapifyByUpToDown() {
	start := 1
	for {
		minIndex := start
		left := start << 1
		right := left + 1
		leftFlag := false
		rightFlag := false
		if left <= p.count && p.nodes[left].dist < p.nodes[start].dist {
			leftFlag = true
			minIndex = left
		}
		if right <= p.count && p.nodes[right].dist < p.nodes[start].dist {
			rightFlag = true
			minIndex = right
		}
		if minIndex == start {
			break
		}
		if leftFlag && rightFlag {
			minIndex = right
			if p.nodes[left].dist < p.nodes[right].dist {
				minIndex = left
			}
		}
		p.index[p.nodes[start].id] = p.nodes[minIndex].id
		p.index[p.nodes[minIndex].id] = p.nodes[start].id
		p.nodes[start], p.nodes[minIndex] = p.nodes[minIndex], p.nodes[start]
		start = minIndex
	}
}
func (p *PriorityQueue) isEmpty() bool {
	return p.count == 0
}
