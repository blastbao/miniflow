package ds

import (
	"testing"
)

func TestDigraph(t *testing.T) {
	fm := []int{0, 0, 0, 6, 6, 5, 3, 2, 2, 8, 7, 9, 9, 9, 11}
	to := []int{1, 6, 5, 4, 9, 4, 5, 0, 3, 7, 6, 10, 11, 12, 12}
	edgeSize := len(fm)

	g := NewDigraph()
	want, actual := 0, g.GetVSize()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}
	for i := 0; i < edgeSize; i++ {
		g.AddEdge(fm[i], to[i])
	}

	want, actual = 13, g.GetVSize()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}

	rg := g.Reverse()
	want, actual = 13, rg.GetVSize()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}

	// delete
	g.DelEdge(0, 1)
	g.DelV(0)
	want, actual = 12, g.GetVSize()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}

	rg = g.Reverse()
	want, actual = 13, rg.GetVSize()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}
}
