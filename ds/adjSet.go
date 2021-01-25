package ds

type set struct {
	size     int
	vertices []int
	seen     map[int]int
}

func newSet() *set {
	s := set{
		size:     0,
		vertices: nil,
		seen:     make(map[int]int),
	}
	return &s
}

func (s *set) isExists(to int) bool {
	_, exists := s.seen[to]
	return exists
}

func (s *set) add(to int) {
	if s.isExists(to) {
		return
	}
	s.vertices = append(s.vertices, to)
	s.seen[to] = s.size
	s.size++
}

func (s *set) delete(to int) {
	if !s.isExists(to) {
		return
	}
	s.size--
	i := s.seen[to]
	last := s.vertices[s.size]
	s.seen[last] = i
	s.vertices[i] = s.vertices[s.size]
	s.vertices[s.size] = 0
	s.vertices = s.vertices[:s.size]
	delete(s.seen, to)
}
