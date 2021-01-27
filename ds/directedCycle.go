package ds

// DirectedCycle finds cycle in directed graph
type DirectedCycle interface {
	IsDAG() bool
	GetCycle() string
}

type directedCycle struct {
	marked  map[int]bool
	edgeTo  map[int]int
	onStack map[int]bool
	cycle   Stack
}

// NewDirectedCycle create a struct for catching cycle in directed grach
func NewDirectedCycle(g Digraph) DirectedCycle {
	c := directedCycle{
		marked:  make(map[int]bool),
		onStack: make(map[int]bool),
		edgeTo:  make(map[int]int),
		cycle:   nil,
	}
	for _, v := range g.V() {
		if c.marked[v] {
			continue
		}
		c.dfsIter(g, v)
	}
	return &c
}

func (c *directedCycle) IsDAG() bool      { return c.cycle == nil }
func (c *directedCycle) GetCycle() string { return c.cycle.String() }

func (c *directedCycle) dfs(G Digraph, v int) {
	c.marked[v] = true
	c.onStack[v] = true
	for _, w := range G.Children(v) {
		if c.cycle != nil {
			return
		}
		if !c.marked[w] {
			c.edgeTo[w] = v
			c.dfs(G, w)
		} else if c.onStack[w] {
			c.cycle = c.catchCycle(v, w)
		}
	}
	c.onStack[v] = false
}

func (c *directedCycle) dfsIter(G Digraph, v int) {
	s := newStack()
	s.Push(v)
	for !s.Empty() && c.cycle == nil {
		w, _ := s.Peek()
		c.marked[w] = true
		c.onStack[w] = true
		i := 0
		for _, x := range G.Children(w) {
			if !c.marked[x] {
				c.edgeTo[x] = w
				s.Push(x)
				i++
			} else if c.onStack[x] {
				c.cycle = c.catchCycle(w, x)
			}
		}
		if i == 0 {
			x, _ := s.Pop()
			c.onStack[x] = false
		}
	}
}

func (c *directedCycle) catchCycle(v, w int) Stack {
	cycle := newStack()
	for x := v; x != w; x = c.edgeTo[x] {
		cycle.Push(x)
	}
	cycle.Push(w)
	cycle.Push(v)
	return cycle
}
