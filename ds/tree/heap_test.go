package tree

import "testing"

type heapDummyItem struct {
	id    int
	key   int
	index int
}

func newHeapDummy(k int) *heapDummyItem {
	return &heapDummyItem{id: 0, key: k, index: -1}
}

func (hd *heapDummyItem) GetID() int         { return hd.id }
func (hd *heapDummyItem) HeapKey() int       { return hd.key }
func (hd *heapDummyItem) HeapIndex() int     { return hd.index }
func (hd *heapDummyItem) SetHeapKey(k int)   { hd.key = k }
func (hd *heapDummyItem) SetHeapIndex(i int) { hd.index = i }

func TestHeap(t *testing.T) {
	keys := []int{2, 0, 6, 4, 8, 10, 12, 14, 16, 18, 20}

	h := &heap{
		items: make([]HeapItem, 3),
		size:  0,
	}

	if !h.Empty() || h.Size() != 0 {
		t.Fatalf("heap size error")
	}

	for i, k := range keys {
		item := newHeapDummy(k)
		h.Add(item)
		t.Logf("%v", cap(h.items))
		want, actual := i+1, h.Size()
		if h.Empty() || h.Size() != i+1 {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}

	for i := 0; i < len(keys); i++ {
		want, actual := i*2, h.Get(1).HeapKey()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}

		want, actual = 1, h.Get(1).HeapIndex()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}

		item := h.Del(1)
		t.Logf("%v", cap(h.items))
		want, actual = i*2, item.HeapKey()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}

		want, actual = -1, item.HeapIndex()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}

	for i, k := range keys {
		item := newHeapDummy(k)
		h.Add(item)
		want, actual := i+1, h.Size()
		if h.Empty() || h.Size() != i+1 {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}

	for i := 0; i < len(keys)-1; i++ {
		j := i + 99
		k := (i + 1) * 2

		// update/set
		h.Set(1, j)
		want, actual := k, h.Get(1).HeapKey()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}

		want, actual = 1, h.Get(1).HeapIndex()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}
}
