package core

type addEdge func(from int, to int)
type addV func(id int)

// Task gets item's info
type Task interface {
	GetID() int
	GetCmd() string
	HeapKey() int
	HeapIndex() int
	SetHeapKey(k int)
	SetHeapIndex(i int)
}

// Item is a vertex in a dag
type Item struct {
	ID         int    `json:"id"`
	Cmd        string `json:"cmd"`
	Upstream   []int  `json:"upstream"`
	Downstream []int  `json:"downstream"`
	priority   int
	index      int
}

// GetID returns item's id
func (it *Item) GetID() int { return it.ID }

// GetCmd returns item's command
func (it *Item) GetCmd() string { return it.Cmd }

// HeapKey return item's priority
func (it *Item) HeapKey() int { return it.priority }

// HeapIndex return item's priority index
func (it *Item) HeapIndex() int { return it.index }

// SetHeapKey update item's priority
func (it *Item) SetHeapKey(k int) { it.priority = k }

// SetHeapIndex update item's index in priority queue
func (it *Item) SetHeapIndex(i int) { it.index = i }

func (it *Item) hasUpstream() bool   { return len(it.Upstream) > 0 }
func (it *Item) hasDownstream() bool { return len(it.Downstream) > 0 }

func (it *Item) process(v addV, e addEdge) {
	v(it.ID)
	it.processUpstream(e)
	it.processDownstream(e)
	it.Upstream = nil
	it.Downstream = nil
}

func (it *Item) processUpstream(e addEdge) {
	if !it.hasUpstream() {
		return
	}
	from := it.ID
	for _, to := range it.Upstream {
		e(from, to)
	}
}

func (it *Item) processDownstream(e addEdge) {
	if !it.hasDownstream() {
		return
	}
	to := it.ID
	for _, from := range it.Downstream {
		e(from, to)
	}
}
