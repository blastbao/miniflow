package basics

import (
	"testing"
)

func TestSet(t *testing.T) {
	testSize := 100000
	s := NewSet()
	want, actual := true, s.Empty()
	if want != actual {
		t.Fatalf("want: %t, actual: %t, want, actual", want, actual)
	}

	for i := 0; i < testSize; i++ {
		s.Add(i)
		want, actual := true, s.Has(i)
		if want != actual {
			t.Fatalf("want: %t, actual: %t", want, actual)
		}
		want, actual = false, s.Empty()
		if want != actual {
			t.Fatalf("want: %t, actual: %t", want, actual)
		}
	}

	for i := 0; i < testSize; i++ {
		s.Add(i)
		want, actual := testSize, s.Size()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}

	for i := 0; i < testSize; i++ {
		s.Del(i)
		want, actual := testSize-i-1, s.Size()
		if want != actual {
			t.Fatalf("want: %d, actual: %d", want, actual)
		}
	}
}
