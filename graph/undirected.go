package graph

import (
	"container/list"
	"fmt"
)

var (
	found bool
)

type UndGraph struct {
	n int
	v []*list.List
}

//初始化一个无向图
func NewUndGraph(n int) *UndGraph {
	v := make([]*list.List, n, n)
	for i := 0; i < n; i++ {
		v[i] = list.New()
	}
	return &UndGraph{
		n: n,
		v: v,
	}
}

func (g *UndGraph) AddEdge(s, t int) {
	g.v[s].PushBack(t)
	g.v[t].PushBack(s)
}

//广度优先
func (g *UndGraph) Bfs(start, end int) []int {
	visited := make([]bool, g.n, g.n)
	visited[start] = true
	visQueue := list.New()
	visQueue.PushBack(start)
	pre := make([]int, g.n, g.n)
	pre[0] = -1
	for visQueue.Len() != 0 {
		cur := visQueue.Front()
		visQueue.Remove(cur)
		curVal := cur.Value.(int)
		for e := g.v[curVal].Front(); e != nil; e = e.Next() {
			d := e.Value.(int)
			if !visited[d] {
				pre[d] = curVal
				if d == end {
					fmt.Println(pre, start, end)
					return pre
				}
				visited[d] = true
				visQueue.PushBack(d)
			}
		}
	}
	return pre
}

//深度优先
func (g *UndGraph) Dfs(start, end int) []int {
	visited := make([]bool, g.n, g.n)
	pre := make([]int, g.n, g.n)
	g.recurDfsItem(start, end, visited, pre)
	return pre
}

func (g *UndGraph) recurDfsItem(start, end int, visited []bool, pre []int) {
	if found == true {
		return
	}
	visited[start] = true
	if start == end {
		found = true
		return
	}
	for cur := g.v[start].Front(); cur != nil; cur = cur.Next() {
		curVal := cur.Value.(int)
		if !visited[curVal] {
			pre[curVal] = start
			g.recurDfsItem(curVal, end, visited, pre)
		}
	}

}

func (g *UndGraph) Print(pre []int, start, end int) {
	if start != end {
		g.Print(pre, start, pre[end])
	}
	fmt.Print(end, " ")
}
