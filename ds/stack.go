package ds

import (
	"errors"
	"fmt"
	"strings"
)

// Stack returns items in filo order
type Stack interface {
	Empty() bool
	Size() int
	Push(v int)
	Pop() (int, error)
	Peek() (int, error)
	String() string
}

type stack struct {
	first *node
	size  int
}

func newStack() Stack {
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

func (s *stack) Pop() (int, error) {
	if s.first == nil {
		return -1, errors.New("stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	s.size--
	return item, nil
}

func (s *stack) Peek() (int, error) {
	if s.first == nil {
		return -1, errors.New("stack is empty")
	}
	return s.first.item, nil
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
