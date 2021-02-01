package core

type roots struct {
	items []int
	size  int
}

func newRoots() *roots {
	return &roots{
		items: make([]int, 2),
		size:  0,
	}
}

func (r *roots) empty() bool { return r.size == 0 }

func (r *roots) add(root int) {
	r.grow()
	r.items[r.size] = root
	r.size++
}

func (r *roots) get() (int, bool) {
	if r.empty() {
		return 0, false
	}
	return r.items[0], true
}

func (r *roots) del() {
	if r.empty() {
		return
	}
	last := r.size - 1
	r.items[0] = r.items[last]
	r.items[last] = 0
	r.size--
}

func (r *roots) grow() {
	if r.size < cap(r.items) {
		return
	}
	items := make([]int, cap(r.items)*2)
	i := 0
	for i < r.size {
		items[i] = r.items[i]
		i++
	}
	r.items = items
}
