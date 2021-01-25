package ds

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := newStack()
	var want, size int
	for i := 0; i <= 10; i++ {
		want, size = i, s.getSize()
		if want != size {
			t.Fatalf("want: %d, actual: %d\n", want, size)
		}
		s.push(i)
	}

	for j := 10; j >= 0; j-- {
		item, err := s.pop()
		if err != nil {
			t.Fatalf(err.Error())
		}
		if item != j {
			t.Fatalf("want: %d, actual: %d\n", j, item)
		}
	}

	_, err := s.pop()
	if err == nil {
		t.Fatalf("want: err(stack empty), actual: %v", err)
	}
}
