package ds

import (
	"errors"
	"fmt"
	"strings"
)

type stack struct {
	first *node
	size  int
}

func newStack() *stack {
	s := stack{nil, 0}
	return &s
}

func (s *stack) isEmpty() bool { return s.first == nil }
func (s *stack) getSize() int  { return s.size }

func (s *stack) push(v int) {
	oldFirst := s.first
	s.first = newNode(v, oldFirst)
	s.size++
}

func (s *stack) pop() (int, error) {
	if s.first == nil {
		return -1, errors.New("stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	s.size--
	return item, nil
}

func (s *stack) String() string {
	if s.first == nil {
		return ""
	}
	var str strings.Builder
	str.Grow(s.size * 4)
	for x := s.first; x != nil; x = x.next {
		fmt.Fprintf(&str, "%d ", x.item)
	}
	return str.String()[:str.Len()-1]
}
