package basics

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	l := &list{
		items: make([]int, 2),
		size:  0,
	}

	testSize := 100000

	for i := 1; i <= testSize; i++ {
		l.Add(i)
		actual, _ := l.Get(i - 1)
		if actual != i {
			t.Fatalf("want: %d, actual: %d", i, actual)
		}
	}
	if l.Size() != testSize {
		t.Fatalf("want: %d, actual: %d", testSize, l.Size())
	}

	for i := 0; i < testSize; i++ {
		c := cap(l.items)
		l.Del(0)
		if l.Size() <= c/4 && cap(l.items) != c/2 {
			t.Fatalf("want: %d, actual: %d", c/2, cap(l.items))
		}
	}

	for i := 1; i <= testSize; i++ {
		l.Add(i)
		actual, _ := l.Get(i - 1)
		if actual != i {
			t.Fatalf("want: %d, actual: %d", i, actual)
		}
	}
	if l.Size() != testSize {
		t.Fatalf("want: %d, actual: %d", testSize, l.Size())
	}
}
