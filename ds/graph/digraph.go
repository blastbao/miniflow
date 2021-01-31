package graph

import (
	"errors"
	"fmt"
	"miniflow/ds/basics"
)

// Digraph provides some directed graph api
type Digraph interface {
	Size() int
	V() []int
	Parents(v int) []int
	Children(v int) []int
	Indegree(v int) int
	Outdegree(v int) int
	AddV(v int)
	AddEdge(v, w int)
	DelV(v int) error
	DelEdge(v, w int) error
	Index(v int) int
}

type adj map[int]Vertex

type digraph struct {
	vertices basics.ArrayList
	adj      adj
}

// NewDigraph create a new directed graph
func NewDigraph() Digraph {
	return &digraph{
		vertices: basics.NewArrayList(),
		adj:      make(adj),
	}
}

func (g *digraph) Index(v int) int {
	if vtx, err := g.getAdj(v); err == nil {
		return vtx.Index()
	}
	return -1
}

// GetVSize return number of vertices
func (g *digraph) Size() int { return g.vertices.Size() }

// GetV return vertices of graph
func (g *digraph) V() []int { return g.vertices.Items() }

// Parents return parents of v
func (g *digraph) Parents(v int) []int {
	if vtx, err := g.getAdj(v); err == nil {
		return vtx.Parents()
	}
	return nil
}

// GetAdj return children of v
func (g *digraph) Children(v int) []int {
	if vtx, err := g.getAdj(v); err == nil {
		return vtx.Children()
	}
	return nil
}

func (g *digraph) Indegree(v int) int {
	if vtx, err := g.getAdj(v); err == nil {
		return vtx.Indegree()
	}
	return 0
}

// GetAdjSize return outdegree of v
func (g *digraph) Outdegree(v int) int {
	if vtx, err := g.getAdj(v); err == nil {
		return vtx.Outdegree()
	}
	return 0
}

//AddV add vertex to graph
func (g *digraph) AddV(v int) { g.addV(v) }

// AddEdge add directed edge (v -> w)
func (g *digraph) AddEdge(v, w int) {
	g.addV(v)
	g.addV(w)
	from, _ := g.getAdj(v)
	from.AddChild(w)
	to, _ := g.getAdj(w)
	to.AddParent(v)
}

// DelV delete v from graph
func (g *digraph) DelV(v int) error {
	vtx, err := g.getAdj(v)
	if err != nil {
		return err
	}
	index := vtx.Index()
	g.delVertex(index)
	for _, child := range vtx.Children() {
		if err = g.delParent(v, child); err != nil {
			return err
		}
	}
	for _, parent := range vtx.Parents() {
		if err = g.delChild(parent, v); err != nil {
			return err
		}
	}
	delete(g.adj, v)
	return nil
}

// DelEdge delete edge (v -> w)
func (g *digraph) DelEdge(v, w int) error {
	if err := g.delParent(v, w); err != nil {
		return err
	}
	if err := g.delChild(v, w); err != nil {
		return err
	}
	return nil
}

func (g *digraph) delVertex(index int) {
	lastIdx := g.vertices.Size() - 1
	last, ok := g.vertices.Get(lastIdx)
	if !ok {
		return
	}
	g.vertices.Del(index)
	g.adj[last].SetIndex(index)
}

func (g *digraph) delParent(from, to int) error {
	child, err := g.getAdj(to)
	if err != nil {
		return err
	}
	child.DelParent(from)
	return nil
}

func (g *digraph) delChild(from, to int) error {
	parent, err := g.getAdj(from)
	if err != nil {
		return err
	}
	parent.DelChild(to)
	return nil
}

func (g *digraph) addV(v int) {
	if _, exists := g.adj[v]; exists {
		return
	}
	g.adj[v] = newVertex(g.Size())
	g.vertices.Add(v)
}

func (g *digraph) getAdj(v int) (Vertex, error) {
	if _, exists := g.adj[v]; !exists {
		err := fmt.Sprintf("vertex not exists: %d\n", v)
		return nil, errors.New(err)
	}
	return g.adj[v], nil
}

func (g digraph) String() string { return fmt.Sprintf("%v", g.adj) }
