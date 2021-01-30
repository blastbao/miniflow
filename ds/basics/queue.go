package basics

// Queue returns items in fifo ordering
type Queue interface {
	Empty() bool
	Size() int
	Items() []int
	Enqueue(v int)
	Dequeue() (int, bool)
}

type queue struct {
	first *node
	last  *node
	size  int
}

// NewQueue creates a queue
func NewQueue() Queue {
	return &queue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *queue) Empty() bool { return q.size == 0 }
func (q *queue) Size() int   { return q.size }

func (q *queue) Items() []int {
	if q.Empty() {
		return nil
	}
	items := make([]int, q.size, q.size)
	cur := q.first
	for i := 0; i < q.size; i++ {
		items[i] = cur.item
		cur = cur.next
	}
	return items
}

func (q *queue) Enqueue(v int) {
	oldLast := q.last
	q.last = newNode(v, nil)
	if q.first == nil {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.size++
}

func (q *queue) Dequeue() (int, bool) {
	if q.first == nil {
		return 0, false
	}
	item := q.first.item
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	q.size--
	return item, true
}
