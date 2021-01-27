package ds

import (
	"testing"
)

func TestDigraph(t *testing.T) {
	fm := []int{0, 0, 0, 6, 6, 5, 3, 2, 2, 8, 7, 9, 9, 9, 11}
	to := []int{1, 6, 5, 4, 9, 4, 5, 0, 3, 7, 6, 10, 11, 12, 12}
	//size := 13

	g := NewDigraph()
	want, actual := 0, g.Size()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}
	for i := 0; i < 15; i++ {
		g.AddEdge(fm[i], to[i])
	}

	tests := []struct {
		want   int
		vertex int
	}{
		{3, 0},
		{0, 1},
		{2, 2},
		{1, 3},
		{0, 4},
		{1, 5},
		{2, 6},
		{1, 7},
		{1, 8},
		{3, 9},
		{0, 10},
		{1, 11},
		{0, 12},
	}

	for _, c := range tests {
		actual = g.Outdegree(c.vertex)
		if c.want != actual {
			t.Errorf("want: %d, actual %d", c.want, actual)
		}
	}

	tests = []struct {
		want   int
		vertex int
	}{
		{1, 0},
		{1, 1},
		{0, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{1, 7},
		{0, 8},
		{1, 9},
		{1, 10},
		{1, 11},
		{2, 12},
	}

	for _, c := range tests {
		actual = g.Indegree(c.vertex)
		if c.want != actual {
			t.Errorf("want: %d, actual %d", c.want, actual)
		}
	}

	tests = []struct {
		want   int
		vertex int
	}{
		{0, 0},
		{1, 1},
		{7, 2},
		{6, 3},
		{4, 4},
		{3, 5},
		{2, 6},
		{9, 7},
		{8, 8},
		{5, 9},
		{10, 10},
		{11, 11},
		{12, 12},
	}

	for _, c := range tests {
		actual = g.Index(c.vertex)
		if c.want != actual {
			t.Errorf("want: %d, actual %d", c.want, actual)
		}
	}

	// delete
	g.DelEdge(0, 1)
	g.DelV(0)
	want, actual = 12, g.Size()
	if want != actual {
		t.Fatalf("want: %d, actual: %d\n", want, actual)
	}
}
