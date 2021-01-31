package core

import (
	"log"
	"miniflow/ds/basics"
	"miniflow/ds/graph"
	"miniflow/ds/tree"
)

// DAG for tasks
type DAG interface {
	Start()
	Produce() <-chan Task
}

type dag struct {
	name    string
	g       graph.Digraph
	pq      tree.MinPQ
	meta    map[int]Task
	maxFlow int
	curFlow int
	roots   basics.ArrayList
	pipeOut chan Task
	success <-chan int
	fail    <-chan int
}

// NewDAG create new dag
func NewDAG(c *Configs, success <-chan int, fail <-chan int) DAG {
	g, meta := initGraph(c)
	d := dag{
		name:    c.Name,
		g:       g,
		meta:    meta,
		maxFlow: c.Parallel,
		curFlow: 0,
		roots:   basics.NewArrayList(),
		pipeOut: make(chan Task),
		success: success,
		fail:    fail,
	}
	return &d
}

func (d *dag) Produce() <-chan Task { return d.pipeOut }

func (d *dag) Start() {
	log.Print("miniflow start")
	log.Print("name: ", d.name)
	log.Print("size: ", d.g.Size())
	if d.hasCycle() {
		return
	}
	d.initPQ()
	d.getSources()
	for d.hasTasks() {

		next, pipeOut := d.switchPipeOut()

		select {
		case id := <-d.fail:
			d.remove(id)
			d.curFlow--

		case id := <-d.success:
			d.settle(id)
			d.curFlow--
			d.getSources()

		case pipeOut <- next:
			d.roots.Del(0)
			d.curFlow++
		}
	}
}

func initGraph(c *Configs) (graph.Digraph, map[int]Task) {
	g := graph.NewDigraph()
	meta := make(map[int]Task, len(c.Tasks))
	for _, task := range c.Tasks {
		task.process(g.AddV, g.AddEdge)
		meta[task.ID] = task
	}
	return g, meta
}

func (d *dag) initPQ() {
	log.Print("dag confirmed, init minPQ")
	vertices := d.g.V()
	if len(vertices) == 0 {
		return
	}
	heapItems := make([]tree.HeapItem, len(vertices)+1)
	i := 1
	for _, v := range vertices {
		task := d.getTaskMeta(v)
		prio := d.g.Outdegree(v)
		task.SetHeapKey(prio)
		heapItems[i] = task
		task.SetHeapIndex(i)
		i++
	}
	d.pq = tree.NewMinPQ(heapItems)
}

func (d *dag) getTaskMeta(v int) Task {
	task, exists := d.meta[v]
	if !exists {
		log.Fatalf("task%d not exists, please check json\n", v)
	}
	return task

}

func (d *dag) hasCycle() bool {
	log.Print("check for directed cycle")
	c := graph.NewDirectedCycle(d.g)
	if !c.IsDAG() {
		log.Print("miniflow has cycle: ", c.GetCycle())
	}
	return !c.IsDAG()
}

func (d *dag) hasTasks() bool { return d.g.Size() > 0 }

func (d *dag) switchPipeOut() (Task, chan Task) {
	var next Task
	var pipeOut chan Task
	if !d.roots.Empty() && d.curFlow < d.maxFlow {
		id, _ := d.roots.Get(0)
		next = d.meta[id]
		pipeOut = d.pipeOut
	}
	return next, pipeOut
}

func (d *dag) getSources() {
	for {
		if d.pq.Empty() || d.pq.GetMin().HeapKey() > 0 {
			break
		}
		root := d.pq.Dequeue()
		d.roots.Add(root.GetID())
	}
}

func (d *dag) remove(id int) {
	topo := graph.NewTopo(d.g, d.g.Parents, id)
	downstream := topo.GetTopoOrder()
	log.Printf("task%d downstream: %v\n", id, downstream[1:])
	postOrder := topo.GetPostOrder()
	for _, id := range postOrder {
		log.Printf("task%d removed\n", id)
		task := d.meta[id]
		d.pq.Remove(task.HeapIndex())
		d.g.DelV(id)
	}
}

func (d *dag) settle(id int) {
	log.Printf("task%d done\n", id)
	d.relaxUpstream(id)
	d.g.DelV(id)
}

func (d *dag) relaxUpstream(upstreamID int) {
	if d.g.Indegree(upstreamID) == 0 {
		return
	}
	downstream := d.g.Parents(upstreamID)
	for _, id := range downstream {
		task := d.meta[id]
		idx := task.HeapIndex()
		prio := d.g.Outdegree(id) - 1
		d.pq.Update(idx, prio)
	}
}
