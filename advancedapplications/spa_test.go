package advancedapplications

import "testing"

func TestSpaGraph_Dijkstra(t *testing.T) {
	s := NewSpaGraph(6)
	s.AddEdge(0, 1, 10)
	s.AddEdge(0, 4, 15)
	s.AddEdge(1, 2, 15)
	s.AddEdge(1, 3, 2)
	s.AddEdge(3, 2, 1)
	s.AddEdge(3, 5, 12)
	s.AddEdge(2, 5, 5)
	s.AddEdge(4, 5, 10)
	s.Dijkstra(0, 5)//0->1->3->2->5
}
