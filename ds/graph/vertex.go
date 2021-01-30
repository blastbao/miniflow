package graph

import "miniflow/ds/basics"

// Vertex is a node in a directed graph
type Vertex interface {
	Indegree() int
	Outdegree() int
	IsSource() bool
	Index() int
	SetIndex(i int)
	Parents() []int
	Children() []int
	AddParent(w int)
	AddChild(w int)
	DelParent(w int)
	DelChild(w int)
}

type vertex struct {
	indegree  basics.Set
	outdegree basics.Set
	index     int
}

func newVertex(i int) Vertex {
	return &vertex{
		indegree:  basics.NewSet(),
		outdegree: basics.NewSet(),
		index:     i,
	}
}

func (vtx *vertex) Indegree() int   { return vtx.indegree.Size() }
func (vtx *vertex) Outdegree() int  { return vtx.outdegree.Size() }
func (vtx *vertex) IsSource() bool  { return vtx.indegree.Empty() }
func (vtx *vertex) Index() int      { return vtx.index }
func (vtx *vertex) SetIndex(i int)  { vtx.index = i }
func (vtx *vertex) Parents() []int  { return vtx.indegree.Items() }
func (vtx *vertex) Children() []int { return vtx.outdegree.Items() }

func (vtx *vertex) AddParent(w int) { vtx.indegree.Add(w) }
func (vtx *vertex) AddChild(w int)  { vtx.outdegree.Add(w) }
func (vtx *vertex) DelParent(w int) { vtx.indegree.Del(w) }
func (vtx *vertex) DelChild(w int)  { vtx.outdegree.Del(w) }
