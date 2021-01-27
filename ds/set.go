package ds

import "fmt"

// Set is a container of non duplicate items
type Set interface {
	Size() int
	IsEmpty() bool
	Items() []int
	Contains(item int) bool
	Add(item int)
	Del(item int)
}

type set struct {
	items []int
	seen  map[int]int
	size  int
}

func newSet() Set {
	s := set{
		items: nil,
		seen:  make(map[int]int),
		size:  0,
	}
	return &s
}

func (s *set) Size() int     { return s.size }
func (s *set) IsEmpty() bool { return s.size == 0 }
func (s *set) Items() []int  { return s.items }
func (s *set) Contains(item int) bool {
	_, ok := s.seen[item]
	return ok
}

func (s *set) Add(item int) {
	if s.Contains(item) {
		return
	}
	s.seen[item] = s.size
	s.size++
	s.items = append(s.items, item)
}

func (s *set) Del(item int) {
	if !s.Contains(item) {
		return
	}
	s.size--
	index := s.seen[item]
	last := s.items[s.size]
	// delete from items
	s.items[index] = last
	// update last item index
	s.seen[last] = index
	s.items[s.size] = 0
	s.items = s.items[:s.size]
	// delete from seen
	delete(s.seen, item)
}

func (s set) String() string { return fmt.Sprintf("%v", s.items) }
