package main

import "fmt"

type IRunner interface {
	Register(ts []*Task)
	Run()
}

type Runner struct {
	chStatus chan string
	Tasks    []*Task
}

func (r *Runner) Register(ts []*Task) {
	r.Tasks = append(r.Tasks, ts...)
}

func (r *Runner) Run() {
	for _, task := range r.Tasks {
		go task.Run(r.chStatus)
	}
	for val := range r.chStatus {
		fmt.Println(val)
	}
}
