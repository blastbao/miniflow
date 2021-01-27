package core

import (
	"log"
	"miniflow/ds"
)

// DAG for tasks
type DAG interface {
	Start()
	Produce() <-chan Task
}

type dag struct {
	name    string
	g       ds.Digraph
	pq      ds.MinPQ
	meta    map[int]Task
	maxFlow int
	curFlow int
	roots   []int
	pipeOut chan Task
	success <-chan int
	fail    <-chan int
}

// NewDAG create new dag
func NewDAG(c *Configs, success <-chan int, fail <-chan int) DAG {
	flowName, g, meta := initGraph(c)
	d := dag{
		name:    flowName,
		g:       g,
		pq:      ds.NewMinPQ(),
		meta:    meta,
		maxFlow: c.Parallel,
		curFlow: 0,
		roots:   nil,
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
			d.roots = d.roots[1:]
			d.curFlow++
		}
	}
}

func initGraph(c *Configs) (string, ds.Digraph, map[int]Task) {
	g := ds.NewDigraph()
	meta := make(map[int]Task, len(c.Tasks))
	for _, task := range c.Tasks {
		task.process(g.AddV, g.AddEdge)
		meta[task.ID] = task
	}
	return c.Name, g, meta
}

func (d *dag) initPQ() {
	log.Print("dag confirmed, init minPQ")
	vertices := d.g.V()
	if len(vertices) == 0 {
		return
	}
	for _, v := range vertices {
		task := d.getTaskMeta(v)
		prio := d.g.Outdegree(v)
		task.SetPrio(prio)
		task.SetIndex(-1)
		d.pq.Enqueue(task)
	}
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
	c := ds.NewDirectedCycle(d.g)
	if !c.IsDAG() {
		log.Print("miniflow has cycle: ", c.GetCycle())
	}
	return !c.IsDAG()
}

func (d *dag) hasTasks() bool { return d.g.Size() > 0 }

func (d *dag) switchPipeOut() (Task, chan Task) {
	var next Task
	var pipeOut chan Task
	if len(d.roots) > 0 && d.curFlow < d.maxFlow {
		next = d.meta[d.roots[0]]
		pipeOut = d.pipeOut
	}
	return next, pipeOut
}

func (d *dag) getSources() {
	for {
		if d.pq.IsEmpty() || d.pq.GetMin().GetPrio() > 0 {
			break
		}
		root := d.pq.Dequeue()
		d.roots = append(d.roots, root.GetID())
	}
}

func (d *dag) remove(id int) {
	topo := ds.NewTopo(d.g, d.g.Parents, id)
	downstream := topo.GetTopoOrder()
	log.Printf("task%d downstream: %v\n", id, downstream[1:])
	postOrder := topo.GetPostOrder()
	for _, id := range postOrder {
		log.Printf("task%d removed\n", id)
		task := d.meta[id]
		d.pq.Remove(task.GetIndex())
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
		idx := task.GetIndex()
		prio := d.g.Outdegree(id) - 1
		d.pq.Update(idx, prio)
	}
}
