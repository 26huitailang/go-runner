package main

import (
	"fmt"
	"sync"
)

type IRunner interface {
	Register(ts []ITask)
	Run()
}

type Runner struct {
	chStatus chan string
	Tasks    []ITask
}

func (r *Runner) Register(ts []ITask) {
	r.Tasks = append(r.Tasks, ts...)
}

func (r *Runner) Run() {
	var wg sync.WaitGroup
	wg.Add(len(r.Tasks) + 1)
	for _, task := range r.Tasks {
		go task.Run(r.chStatus, &wg)
	}
	go func(c chan string) {
		wg.Wait()
		close(c)
	}(r.chStatus)
	for val := range r.chStatus {
		fmt.Println(val)
	}
	wg.Done()
}
