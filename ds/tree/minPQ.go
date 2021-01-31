package tree

// MinPQ implements min priority queue
type MinPQ interface {
	Empty() bool
	Enqueue(item HeapItem)
	Dequeue() HeapItem
	GetMin() HeapItem
	Update(i, prio int)
	Remove(i int)
}

type minPQ struct{ h Heap }

// NewMinPQ creates a new min priority queue
func NewMinPQ(items []HeapItem) MinPQ {
	return &minPQ{h: NewHeap(items)}
}

func (m *minPQ) Empty() bool           { return m.h.Empty() }
func (m *minPQ) Enqueue(item HeapItem) { m.h.Add(item) }
func (m *minPQ) Dequeue() HeapItem     { return m.h.Del(1) }
func (m *minPQ) GetMin() HeapItem      { return m.h.Get(1) }
func (m *minPQ) Update(i, prio int)    { m.h.Set(i, prio) }
func (m *minPQ) Remove(i int)          { m.h.Del(i) }
