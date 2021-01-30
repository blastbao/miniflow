package graph

import (
	"miniflow/ds/basics"
)

type order func(v int) []int

// Topo gives topological order
type Topo interface {
	GetTopoOrder() []int
	GetPostOrder() []int
}

type topo struct {
	marked      map[int]bool
	reversePost basics.Stack
	post        basics.Queue
}

// NewTopo creates new Topo interface
func NewTopo(g Digraph, di order, v int) Topo {
	t := topo{
		marked:      make(map[int]bool),
		reversePost: basics.NewStack(),
		post:        basics.NewQueue(),
	}
	t.dfsIter(g, di, v)
	return &t
}

func (t *topo) GetTopoOrder() []int { return t.reversePost.Items() }
func (t *topo) GetPostOrder() []int { return t.post.Items() }

func (t *topo) dfs(G Digraph, di order, v int) {
	t.marked[v] = true
	for _, w := range di(v) {
		if t.marked[w] {
			continue
		}
		t.dfs(G, di, w)
	}
	t.post.Enqueue(v)
	t.reversePost.Push(v)
}

func (t *topo) dfsIter(G Digraph, di order, v int) {
	s := basics.NewStack()
	s.Push(v)
	for !s.Empty() {
		w, _ := s.Peek()
		t.marked[w] = true
		i := 0
		for _, x := range di(w) {
			if t.marked[x] {
				continue
			}
			s.Push(x)
			i++
		}
		if i == 0 {
			x, _ := s.Pop()
			t.post.Enqueue(x)
			t.reversePost.Push(x)
		}
	}
}
