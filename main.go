package main

import (
	"log"
	"miniflow/core"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("[usage]: ./miniflow /path/to/tasks.json")
	}

	success := make(chan int)
	fail := make(chan int)
	conf := core.NewConfigs(os.Args[1])

	dag := core.NewDAG(conf, success, fail)
	executor := core.NewExecutor(dag.Produce(), success, fail)
	executor.Start()
	dag.Start()
}
