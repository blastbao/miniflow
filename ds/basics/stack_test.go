package basics

import (
	"testing"
)

func TestStack(t *testing.T) {
	testSize := 1000000
	s := NewStack()
	var want, size int
	for i := 0; i <= testSize; i++ {
		want, size = i, s.Size()
		if want != size {
			t.Fatalf("want: %d, actual: %d\n", want, size)
		}
		s.Push(i)
	}

	for j := testSize; j >= 0; j-- {
		item, ok := s.Pop()
		if ok != true {
			t.Fatalf("pop failed")
		}
		if item != j {
			t.Fatalf("want: %d, actual: %d\n", j, item)
		}
	}

	_, ok := s.Pop()
	if ok == true {
		t.Fatalf("want: false, actual: %t", ok)
	}
}
