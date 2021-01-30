package basics

import (
	"fmt"
	"strings"
)

// Stack returns items in filo order
type Stack interface {
	Empty() bool
	Size() int
	Push(v int)
	Pop() (int, bool)
	Peek() (int, bool)
	Items() []int
	String() string
}

type stack struct {
	first *node
	size  int
}

// NewStack creates a stack
func NewStack() Stack {
	s := stack{nil, 0}
	return &s
}

func (s *stack) Empty() bool { return s.first == nil }
func (s *stack) Size() int   { return s.size }

func (s *stack) Push(v int) {
	oldFirst := s.first
	s.first = newNode(v, oldFirst)
	s.size++
}

func (s *stack) Pop() (int, bool) {
	if s.first == nil {
		return 0, false
	}
	item := s.first.item
	s.first = s.first.next
	s.size--
	return item, true
}

func (s *stack) Peek() (int, bool) {
	if s.first == nil {
		return 0, false
	}
	return s.first.item, true
}

func (s *stack) Items() []int {
	if s.Empty() {
		return nil
	}
	items := make([]int, s.size, s.size)
	cur := s.first
	for i := 0; i < s.size; i++ {
		items[i] = cur.item
		cur = cur.next
	}
	return items
}

func (s *stack) String() string {
	if s.first == nil {
		return ""
	}
	var str strings.Builder
	for x := s.first; x != nil; x = x.next {
		fmt.Fprintf(&str, "%d->", x.item)
	}
	return str.String()[:str.Len()-2]
}
