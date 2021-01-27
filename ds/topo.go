package ds

type order func(v int) []int

// Topo gives topological order
type Topo interface {
	GetTopoOrder() []int
	GetPostOrder() []int
}

type topo struct {
	marked      map[int]bool
	reversePost Stack
	post        Queue
}

// NewTopo creates new Topo interface
func NewTopo(g Digraph, di order, v int) Topo {
	t := topo{
		marked:      make(map[int]bool),
		reversePost: newStack(),
		post:        newQueue(),
	}
	t.dfsIter(g, di, v)
	return &t
}

func (t *topo) GetTopoOrder() []int {
	var order []int
	for t.reversePost.Size() > 0 {
		v, err := t.reversePost.Pop()
		if err != nil {
			return nil
		}
		order = append(order, v)
	}
	return order
}

func (t *topo) GetPostOrder() []int {
	var order []int
	for !t.post.Empty() {
		v, err := t.post.Dequeue()
		if err != nil {
			return nil
		}
		order = append(order, v)
	}
	return order
}

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
	s := newStack()
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
