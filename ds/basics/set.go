package basics

// Set is a container of non duplicate items
type Set interface {
	Size() int
	Empty() bool
	Items() []int
	Has(item int) bool
	Add(item int)
	Del(item int)
}

type set struct {
	items ArrayList
	seen  map[int]int
}

// NewSet creates a new set
func NewSet() Set {
	s := set{
		items: NewArrayList(),
		seen:  make(map[int]int),
	}
	return &s
}

func (s *set) Size() int    { return s.items.Size() }
func (s *set) Empty() bool  { return s.items.Size() == 0 }
func (s *set) Items() []int { return s.items.Items() }
func (s *set) Has(item int) bool {
	_, ok := s.seen[item]
	return ok
}

func (s *set) Add(item int) {
	if s.Has(item) {
		return
	}
	s.seen[item] = s.items.Size()
	s.items.Add(item)
}

func (s *set) Del(item int) {
	if !s.Has(item) {
		return
	}
	index := s.seen[item]
	last, _ := s.items.Get(s.items.Size() - 1)
	s.items.Del(index)
	delete(s.seen, item)
	// update last item index
	s.seen[last] = index
}
