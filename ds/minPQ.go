package ds

// MinPQItem in the queue
type MinPQItem interface {
	GetID() int
	GetPrio() int
	GetIndex() int
	SetPrio(n int)
	SetIndex(n int)
}

// MinPQ implements min priority queue
type MinPQ interface {
	IsEmpty() bool
	GetMin() MinPQItem
	Enqueue(item MinPQItem)
	Dequeue() MinPQItem
	Update(i, prio int)
	Remove(i int) MinPQItem
}

type minPQ struct {
	pq []MinPQItem
	n  int
}

// NewMinPQ creates a new queue
func NewMinPQ() MinPQ {
	pq := make([]MinPQItem, 2)
	return &minPQ{pq: pq, n: 0}
}

// IsEmpty test if pq is empty
func (pq *minPQ) IsEmpty() bool { return pq.n == 0 }

// Enqueue push item to queue
func (pq *minPQ) Enqueue(item MinPQItem) {
	length := len(pq.pq)
	if pq.n == length-1 {
		pq.resize(length * 2)
	}
	pq.n++
	pq.pq[pq.n] = item
	item.SetIndex(pq.n)
	pq.swim(pq.n)
}

// GetMin get item with the min priority from queue
func (pq *minPQ) GetMin() MinPQItem {
	if pq.n == 0 {
		return nil
	}
	return pq.pq[1]
}

// Dequeue pop the item with min priority
func (pq *minPQ) Dequeue() MinPQItem { return pq.Remove(1) }

// Remove item from pq
func (pq *minPQ) Remove(i int) MinPQItem {
	if pq.n == 0 || i < 1 {
		return nil
	}
	item := pq.pq[i]
	if i > 0 && i < pq.n {
		pq.exch(i, pq.n)
	}
	pq.n--
	pq.pq[pq.n+1] = nil
	item.SetIndex(-1)
	pq.sinkThenResize(i)
	return item
}

// Update item in pq
func (pq *minPQ) Update(i, prio int) {
	if pq.n == 0 && i < 1 {
		return
	}
	if i > 0 && i < pq.n {
		pq.exch(i, pq.n)
		pq.n--
		pq.sink(i)
		pq.n++
	}
	pq.pq[pq.n].SetPrio(prio)
	pq.swim(pq.n)
}

func (pq *minPQ) greater(i, j int) bool { return pq.pq[i].GetPrio() > pq.pq[j].GetPrio() }
func (pq *minPQ) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
	pq.pq[i].SetIndex(i)
	pq.pq[j].SetIndex(j)
}

func (pq *minPQ) sinkThenResize(k int) {
	pq.sink(k)
	length := len(pq.pq)
	if pq.n > 0 && pq.n == (length-1)/4 {
		pq.resize(length / 2)
	}
}

func (pq *minPQ) resize(newSize int) {
	arr := make([]MinPQItem, newSize)
	for i := 1; i <= pq.n; i++ {
		arr[i] = pq.pq[i]
	}
	pq.pq = arr
}

func (pq *minPQ) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}
}

func (pq *minPQ) sink(k int) {
	for k*2 <= pq.n {
		j := k * 2
		if j < pq.n && pq.greater(j, j+1) {
			j++
		}
		if !pq.greater(k, j) {
			break
		}
		pq.exch(j, k)
		k = j
	}
}
