package ds

import (
	"testing"
)

type pqTestItem struct {
	id      int
	prioity int
	index   int
}

func newTestItem(id int, prio int) *pqTestItem {
	return &pqTestItem{id: id, prioity: prio, index: -1}
}

func (it *pqTestItem) GetID() int     { return it.id }
func (it *pqTestItem) GetPrio() int   { return it.prioity }
func (it *pqTestItem) GetIndex() int  { return it.index }
func (it *pqTestItem) SetPrio(n int)  { it.prioity = n }
func (it *pqTestItem) SetIndex(n int) { it.index = n }

func TestMinPQ(t *testing.T) {
	t0 := 2
	t1 := 0
	t2 := 6
	t3 := 4

	items := map[int]int{
		0: t2, 1: t1, 2: t3, 3: t0,
	}

	pq := NewMinPQ()

	for k, priority := range items {
		item := newTestItem(k, priority)
		pq.Enqueue(item)
	}

	for i := 0; i < len(items); i++ {
		// get min
		j := i * 2
		want, actual := j, pq.GetMin().GetPrio()
		if want != actual {
			t.Fatalf("want: %v, actual: %v, priority not in order\n", want, actual)
		}
		// dequeue
		actual = pq.Dequeue().GetPrio()
		if want != actual {
			t.Fatalf("want: %v, actual: %v, priority not in order\n", want, actual)
		}
	}

	for k, priority := range items {
		item := newTestItem(k, priority)
		pq.Enqueue(item)
	}

	for i := 0; i < len(items)-1; i++ {
		j := i + 99
		k := (i + 1) * 2

		// update
		pq.Update(1, j)
		want, actual := k, pq.GetMin().GetPrio()
		if want != actual {
			t.Fatalf("want: %v, actual: %v, priority not in order\n", want, actual)
		}
	}
}
