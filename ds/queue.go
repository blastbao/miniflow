package ds

import (
	"errors"
	"fmt"
	"strings"
)

type queue struct {
	first *node
	last  *node
	size  int
}

func newQueue() *queue {
	return &queue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (q *queue) isEmpty() bool { return q.size == 0 }
func (q *queue) getSize() int  { return q.size }

func (q *queue) enqueue(v int) {
	oldLast := q.last
	q.last = newNode(v, nil)
	if q.first == nil {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.size++
}

func (q *queue) dequeue() (int, error) {
	if q.first == nil {
		return -1, errors.New("queue is empty")
	}
	item := q.first.item
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	q.size--
	return item, nil
}
func (q *queue) String() string {
	if q.first == nil {
		return ""
	}
	var str strings.Builder
	str.Grow(q.size * 4)
	for x := q.first; x != nil; x = x.next {
		fmt.Fprintf(&str, "%d ", x.item)
	}
	return str.String()[:str.Len()-1]
}
