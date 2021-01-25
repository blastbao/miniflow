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
	cycle   *stack
}

// NewDirectedCycle create a struct for catching cycle in directed grach
func NewDirectedCycle(g Digraph) DirectedCycle {
	c := directedCycle{
		marked:  make(map[int]bool),
		onStack: make(map[int]bool),
		edgeTo:  make(map[int]int),
		cycle:   nil,
	}
	for _, v := range g.GetV() {
		if c.marked[v] {
			continue
		}
		c.dfs(g, v)
	}
	return &c
}

func (c *directedCycle) IsDAG() bool      { return c.cycle == nil }
func (c *directedCycle) GetCycle() string { return c.cycle.String() }

func (c *directedCycle) dfs(G Digraph, v int) {
	c.marked[v] = true
	c.onStack[v] = true
	for _, w := range G.GetAdj(v) {
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

func (c *directedCycle) catchCycle(v, w int) *stack {
	cycle := newStack()
	for x := v; x != w; x = c.edgeTo[x] {
		cycle.push(x)
	}
	cycle.push(w)
	cycle.push(v)
	return cycle
}
