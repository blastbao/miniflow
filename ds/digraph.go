package ds

import (
	"errors"
	"fmt"
)

// Digraph provides some directed graph api
type Digraph interface {
	GetVSize() int
	GetV() []int
	GetAdj(v int) []int
	GetAdjSize(v int) int
	AddV(v int)
	AddEdge(v, w int)
	DelV(v int) error
	DelEdge(v, w int) error
	Reverse() Digraph
}

type adj map[int]*set

type digraph struct {
	size     int
	vertices []int
	adj      *adj
	seen     map[int]int
}

// NewDigraph create a new directed graph
func NewDigraph() Digraph {
	adj := make(adj)
	seen := make(map[int]int)
	return &digraph{size: 0, adj: &adj, seen: seen}
}

// GetVSize return number of vertices
func (g *digraph) GetVSize() int { return g.size }

// GetV return vertices of graph
func (g *digraph) GetV() []int { return g.vertices }

// GetAdj return children of v
func (g *digraph) GetAdj(v int) []int {
	set, _ := g.getAdj(v)
	return set.vertices
}

// GetAdjSize return children size of v
func (g *digraph) GetAdjSize(v int) int {
	set, _ := g.getAdj(v)
	return set.size
}

//AddV add vertex to graph
func (g *digraph) AddV(v int) { g.addV(v) }

// AddEdge add directed edge to graph
func (g *digraph) AddEdge(from int, to int) {
	g.addV(from)
	g.addV(to)
	set, _ := g.getAdj(from)
	set.add(to)
}

// DelV delete v from graph
func (g *digraph) DelV(v int) error {
	if _, exists := (*g.adj)[v]; !exists {
		err := fmt.Sprintf("vertex not exists: %d\n", v)
		return errors.New(err)
	}
	g.size--
	i := g.seen[v]
	last := g.vertices[g.size]
	g.seen[last] = i
	g.vertices[i] = g.vertices[g.size]
	g.vertices[g.size] = 0
	g.vertices = g.vertices[:g.size]
	delete(*g.adj, v)
	delete(g.seen, v)
	return nil
}

// DelEdge delete edge (v -> w)
func (g *digraph) DelEdge(v, w int) error {
	set, err := g.getAdj(v)
	if err != nil {
		return err
	}
	set.delete(w)
	return nil
}

// Reverse makes a reverse graph
func (g *digraph) Reverse() Digraph {
	rg := NewDigraph()
	for from := range *g.adj {
		rg.AddV(from)
		set, _ := g.getAdj(from)
		for _, to := range set.vertices {
			rg.AddEdge(to, from)
		}
	}
	return rg
}

func (g *digraph) addV(v int) {
	if _, exists := (*g.adj)[v]; exists {
		return
	}
	(*g.adj)[v] = newSet()
	g.vertices = append(g.vertices, v)
	g.seen[v] = g.size
	g.size++
}

func (g *digraph) getAdj(v int) (*set, error) {
	if _, exists := (*g.adj)[v]; !exists {
		err := fmt.Sprintf("adj not exists: %d\n", v)
		return nil, errors.New(err)
	}
	return (*g.adj)[v], nil
}

func (s set) String() string     { return fmt.Sprintf("%v", s.vertices) }
func (g digraph) String() string { return fmt.Sprintf("%v", g.adj) }
