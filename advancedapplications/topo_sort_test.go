package advancedapplications

import "testing"

func TestGraph_TopoSortByKahn(t *testing.T) {
	g := NewGraph(100)
	g.AddPoint("A")
	g.AddPoint("B")
	g.AddPoint("C")
	g.AddPoint("E")
	g.AddPoint("M")
	g.AddPoint("N")
	g.AddEdge("A", "B")
	g.AddEdge("C", "B")
	g.AddEdge("B", "E")
	g.AddEdge("E", "M")
	g.AddEdge("E", "N")
	g.TopoSortByKahn()
}

func TestGraph_TopoSortByDFS(t *testing.T) {
	g := NewGraph(100)
	g.AddPoint("A")
	g.AddPoint("B")
	g.AddPoint("C")
	g.AddPoint("E")
	g.AddPoint("M")
	g.AddPoint("N")
	g.AddEdge("A", "B")
	g.AddEdge("C", "B")
	g.AddEdge("B", "E")
	g.AddEdge("E", "M")
	g.AddEdge("E", "N")
	g.TopoSortByDFS()
}