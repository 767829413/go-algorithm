package advancedapplications

import "testing"

func TestSpaGraphAStart(t *testing.T) {
	s := NewSpaGraphAStart(11)
	s.AddVetex(0, 320, 630)
	s.AddVetex(1, 300, 630)
	s.AddVetex(2, 280, 625)
	s.AddVetex(3, 270, 630)
	s.AddVetex(4, 320, 700)
	s.AddVetex(5, 360, 620)
	s.AddVetex(6, 320, 590)
	s.AddVetex(7, 370, 580)
	s.AddVetex(8, 350, 730)
	s.AddVetex(9, 390, 690)
	s.AddVetex(10, 400, 620)
	s.AddEdge(0, 1, 20)
	s.AddEdge(1, 2, 20)
	s.AddEdge(2, 3, 10)
	s.AddEdge(0, 4, 60)
	s.AddEdge(4, 8, 50)
	s.AddEdge(8, 9, 50)
	s.AddEdge(9, 10, 60)
	s.AddEdge(0, 5, 60)
	s.AddEdge(5, 8, 70)
	s.AddEdge(5, 9, 80)
	s.AddEdge(5, 10, 50)
	s.AddEdge(0, 6, 60)
	s.AddEdge(6, 7, 70)
	s.AddEdge(7, 10, 110)
	s.AStar(0, 10)
}
