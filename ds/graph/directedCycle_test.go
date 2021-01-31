package graph

import (
	"testing"
)

func TestDirectedCycle(t *testing.T) {
	fm := []int{4, 2, 3}
	to := []int{2, 3, 2}
	//fm := []int{4, 2, 3, 6, 0, 2, 11, 12, 9, 9, 8, 10, 11, 4, 3, 7, 8, 5, 0, 6, 6, 7}
	//to := []int{2, 3, 2, 0, 1, 0, 12, 9, 10, 11, 9, 12, 4, 3, 5, 8, 7, 4, 5, 4, 9, 6}
	edgeSize := len(fm)

	g := NewDigraph()
	for i := 0; i < edgeSize; i++ {
		g.AddEdge(fm[i], to[i])
	}

	dc := NewDirectedCycle(g)
	t.Log(dc.GetCycle())
	if dc.IsDAG() == true {
		t.Fatal("should not be a dag")
	}
}
