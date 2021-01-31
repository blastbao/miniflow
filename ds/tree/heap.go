package tree

// HeapItem is an item in the heap
type HeapItem interface {
	GetID() int
	HeapKey() int
	HeapIndex() int
	SetHeapKey(k int)
	SetHeapIndex(i int)
}

// Heap implements a heap
type Heap interface {
	Empty() bool
	Size() int
	Add(item HeapItem)
	Set(i, k int)
	Get(i int) HeapItem
	Del(i int) HeapItem
}

// NewHeap createa a heap
func NewHeap() Heap {
	h := heap{
		items: make([]HeapItem, 3),
		size:  0,
	}
	return &h
}

type heap struct {
	items []HeapItem
	size  int
}

func (h *heap) Empty() bool { return h.size == 0 }
func (h *heap) Size() int   { return h.size }

func (h *heap) Add(item HeapItem) {
	h.grow()
	h.size++
	h.items[h.size] = item
	item.SetHeapIndex(h.size)
	h.swim(h.size)
}

func (h *heap) Set(i, k int) {
	if h.Empty() || !h.inRange(i) {
		return
	}
	h.items[i].SetHeapKey(k)
	h.swim(i)
	h.sink(i)
}

func (h *heap) Get(i int) HeapItem {
	if h.Empty() || !h.inRange(i) {
		return nil
	}
	return h.items[i]
}

func (h *heap) Del(i int) HeapItem {
	if h.Empty() || !h.inRange(i) {
		return nil
	}
	item := h.items[i]
	if i < h.size {
		h.swap(i, h.size)
	}
	h.items[h.size] = nil
	h.size--
	h.sink(i)
	h.shrink()
	item.SetHeapIndex(-1)
	return item
}

func (h *heap) grow() {
	if h.size+1 < cap(h.items) {
		return
	}
	h.resize((cap(h.items) + 1) * 2)
}

func (h *heap) shrink() {
	if h.Empty() {
		return
	}
	if h.size > (cap(h.items)-1)/4 {
		return
	}
	h.resize(cap(h.items) / 2)
}

func (h *heap) resize(cap int) {
	items := make([]HeapItem, cap)
	i := 1
	for i <= h.size {
		items[i] = h.items[i]
		i++
	}
	h.items = items
}

func (h *heap) inRange(i int) bool { return i >= 1 && i <= h.size }

func (h *heap) compare(i, j int) bool {
	return h.items[i].HeapKey() > h.items[j].HeapKey()
}

func (h *heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].SetHeapIndex(i)
	h.items[j].SetHeapIndex(j)
}

func (h *heap) swim(k int) {
	for k > 1 && h.compare(k/2, k) {
		h.swap(k/2, k)
		k = k / 2
	}
}

func (h *heap) sink(k int) {
	for k*2 <= h.size {
		j := k * 2
		if j < h.size && h.compare(j, j+1) {
			j++
		}
		if !h.compare(k, j) {
			break
		}
		h.swap(j, k)
		k = j
	}
}
