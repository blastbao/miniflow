package basics

// ArrayList implements array resizing
type ArrayList interface {
	Empty() bool
	Size() int
	Items() []int
	Add(item int)
	Get(i int) (int, bool)
	Set(i int, item int)
	Del(i int)
	Swap(i, j int)
}

// NewArrayList creates a new ArrayList
func NewArrayList() ArrayList {
	return &list{
		items: make([]int, 2),
		size:  0,
	}
}

type list struct {
	items []int
	size  int
}

func (l *list) Empty() bool  { return l.size == 0 }
func (l *list) Size() int    { return l.size }
func (l *list) Items() []int { return l.items[:l.size] }

func (l *list) Get(i int) (int, bool) {
	if !l.inRange(i) {
		return 0, false
	}
	return l.items[i], true
}

func (l *list) Set(i int, item int) {
	if !l.inRange(i) {
		return
	}
	l.items[i] = item
}

func (l *list) Del(i int) {
	if l.Empty() || !l.inRange(i) {
		return
	}
	l.size--
	l.items[i] = l.items[l.size]
	l.items[l.size] = 0
	l.shrink()
}

func (l *list) Add(item int) {
	l.grow()
	l.items[l.size] = item
	l.size++
}

func (l *list) Swap(i, j int) {
	if l.Empty() {
		return
	}
	if !l.inRange(i) || !l.inRange(j) || i == j {
		return
	}
	l.items[i], l.items[j] = l.items[j], l.items[i]
}

func (l *list) inRange(i int) bool { return i >= 0 && i < l.size }

func (l *list) shrink() {
	if l.size > cap(l.items)/4 {
		return
	}
	l.resize(cap(l.items) / 2)
}

func (l *list) grow() {
	if l.size+1 <= cap(l.items) {
		return
	}
	l.resize(cap(l.items) * 2)
}

func (l *list) resize(capacity int) {
	items := make([]int, capacity)
	i := 0
	for i < l.size {
		items[i] = l.items[i]
		i++
	}
	l.items = items
}
