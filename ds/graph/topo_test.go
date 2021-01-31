package graph

import (
	"testing"
)

func TestTopoOrder(t *testing.T) {
	fm := []int{0, 0, 0, 6, 6, 5, 3, 2, 2, 8, 7, 9, 9, 9, 11}
	to := []int{1, 6, 5, 4, 9, 4, 5, 0, 3, 7, 6, 10, 11, 12, 12}
	edgeSize := len(fm)

	g := NewDigraph()
	for i := 0; i < edgeSize; i++ {
		g.AddEdge(fm[i], to[i])
	}

	topoIn := NewTopo(g, g.Parents, 5)
	t.Log("indegree order")
	t.Log("post: ", topoIn.GetPostOrder())
	t.Log("topo: ", topoIn.GetTopoOrder())

	topoOut := NewTopo(g, g.Children, 5)
	t.Log("outdegree order")
	t.Log("post: ", topoOut.GetPostOrder())
	t.Log("topo: ", topoOut.GetTopoOrder())
}
