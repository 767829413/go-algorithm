package advancedapplications

import (
	"container/list"
	"fmt"
)

/**
拓扑排序的两种简单实现
1. 构建数据结构(有向无环图)
2. 使用Kahn算法
3. 使用DFS算法
 */

type graph struct {
	v    int                   //顶点个数
	obj  map[string]*list.List //构建邻接表
	robj map[string]*list.List //构建逆邻接表
}

func NewGraph(v int) *graph {
	obj := make(map[string]*list.List, v)
	robj := make(map[string]*list.List, v)
	return &graph{
		v:    v,
		obj:  obj,
		robj: robj,
	}
}

func (g *graph) AddPoint(p string) {
	if _, ok := g.obj[p]; !ok {
		g.obj[p] = list.New()
		g.robj[p] = list.New()
	}
}

//添加边 定点a -> b a要先于b
func (g *graph) AddEdge(a, b string) {
	if _, ok := g.obj[a]; !ok {
		g.obj[a] = list.New()
	}
	if _, ok := g.robj[b]; !ok {
		g.robj[b] = list.New()
	}
	g.obj[a].PushBack(b)
	g.robj[b].PushBack(a)
}

func (g *graph) TopoSortByKahn() {
	inDegree := g.initInDegree()
	// 统计每个顶点的入度
	for _, l := range g.obj {
		if l.Len() != 0 {
			w := l.Front()
			for w != nil {
				inDegree[w.Value.(string)] += 1
				w = w.Next()
			}
		}
	}
	queue := list.New()
	for k, v := range inDegree {
		if v == 0 {
			queue.PushBack(k)
		}
	}
	for queue.Len() != 0 {
		q := queue.Front()
		i := queue.Remove(q).(string)
		fmt.Print("->", i)
		for w := g.obj[i].Front(); w != nil; w = w.Next() {
			cur := w.Value.(string)
			inDegree[cur] -= 1
			if inDegree[cur] == 0 {
				queue.PushBack(cur)
			}
		}
	}
	fmt.Println()
}

func (g *graph) initInDegree() map[string]int {
	inDegree := make(map[string]int, g.v)
	for k := range g.obj {
		inDegree[k] = 0
	}
	return inDegree
}

func (g *graph) TopoSortByDFS() {
	visted := make(map[string]bool, g.v)
	for k := range g.obj {
		if _, ok := visted[k]; !ok {
			visted[k] = true
			g.dfs(k, visted)
		}
	}
}

func (g *graph) dfs(k string, visted map[string]bool) {
	l := g.robj[k]
	if l.Len() != 0 {
		w := l.Front()
		for ; w != nil; w = w.Next() {
			if visted[w.Value.(string)] {
				continue
			}
			visted[w.Value.(string)] = true
			g.dfs(w.Value.(string), visted)
		}
	}
	fmt.Print("->", k)
}
