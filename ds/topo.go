package ds

// Topo gives topological order
type Topo interface {
	GetTopoOrder() []int
	GetPostOrder() []int
}

type topo struct {
	marked      map[int]bool
	reversePost *stack
	post        *queue
}

// NewTopo creates new Topo interface
func NewTopo(g Digraph, v int) Topo {
	t := topo{
		marked:      make(map[int]bool),
		reversePost: newStack(),
		post:        newQueue(),
	}
	t.dfs(g, v)
	return &t
}

func (t *topo) GetTopoOrder() []int {
	var order []int
	for t.reversePost.size > 0 {
		v, err := t.reversePost.pop()
		if err != nil {
			return nil
		}
		order = append(order, v)
	}
	return order
}

func (t *topo) GetPostOrder() []int {
	var order []int
	for !t.post.isEmpty() {
		v, err := t.post.dequeue()
		if err != nil {
			return nil
		}
		order = append(order, v)
	}
	return order
}

func (t *topo) dfs(G Digraph, v int) {
	t.marked[v] = true
	for _, w := range G.GetAdj(v) {
		if t.marked[w] {
			continue
		}
		t.dfs(G, w)
	}
	t.post.enqueue(v)
	t.reversePost.push(v)
}
