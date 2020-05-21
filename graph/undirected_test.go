package graph

import "testing"

func getUndGraph() *UndGraph {
	g := NewUndGraph(20)
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(3, 4)
	g.AddEdge(2, 5)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)
	return g
}

func TestUndGraph_Bfs(t *testing.T) {
	g := getUndGraph()
	pre := g.Bfs(0, 7)
	g.Print(pre, 0, 7)
}

func TestUndGraph_Dfs(t *testing.T) {
	g := getUndGraph()
	pre := g.Dfs(0, 7)
	g.Print(pre, 0, 7)
}
