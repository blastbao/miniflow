package core

import (
	"log"
	"os/exec"
	"strings"
)

// Executor consume tasks
type Executor interface {
	Start()
}

type executor struct {
	pipeIn  <-chan Task
	success chan<- int
	fail    chan<- int
}

// NewExecutor create a new executor interface
func NewExecutor(in <-chan Task, success chan<- int, fail chan<- int) Executor {
	e := executor{
		pipeIn:  in,
		success: success,
		fail:    fail,
	}
	return &e
}

func (e *executor) Start() {
	go e.run()
}

func (e *executor) run() {
	for {
		select {
		case task := <-e.pipeIn:
			e.execute(task)
		}
	}
}

func (e *executor) execute(t Task) {
	go func() {
		id, cmds := t.GetID(), t.GetCmd()
		log.Printf("task%d start : %s\n", id, cmds)
		err := e.execCmds(id, cmds)
		if err != nil {
			log.Printf("task%d failed: %s\n", id, err)
			e.fail <- id
			return
		}
		e.success <- t.GetID()
	}()
}

func (e *executor) execCmds(id int, cmdstr string) error {
	cmds := split(cmdstr, ';')
	for _, cmd := range cmds {
		err := e.execCmd(id, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *executor) execCmd(id int, cmdstr string) error {
	cmd := split(cmdstr, ' ')
	c := exec.Command(cmd[0], cmd[1:]...)
	out, err := c.CombinedOutput()
	if err != nil {
		return err
	}
	outStr := strings.TrimSpace(string(out))
	if outStr != "" {
		log.Printf("task%d output: %s\n", id, outStr)
	}
	return nil
}

func split(s string, r rune) []string {
	f := func(c rune) bool { return c == r }
	return strings.FieldsFunc(s, f)
}
