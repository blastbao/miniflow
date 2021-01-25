package ds

type node struct {
	item int
	next *node
}

func newNode(v int, next *node) *node {
	return &node{item: v, next: next}
}
