package main

import (
	"fmt"
	"sync"
	"time"
)

type ITask interface {
	Run(c chan<- string, wg *sync.WaitGroup)
}

type Task struct {
	Name   string
	Status bool
}

func NewTask(name string, status bool) *Task {
	return &Task{
		Name:   name,
		Status: status,
	}
}

func (t *Task) Run(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%v closed", t.Name)
			c <- msg
		}
	}()
	t.Status = true
	for i := 0; ; i++ {
		msg := fmt.Sprintf("%v %v", t.Name, i)
		c <- msg
		time.Sleep(1 * time.Second)
	}
}

type TaskAdd struct {
	Task
}

func NewTaskAdd() *TaskAdd {
	return &TaskAdd{
		Task{
			Name:   "add",
			Status: false,
		},
	}
}

func (t *TaskAdd) Run(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	t.Status = true
	a, b := 0, 1
	ret := make(chan int, 10)
	go func() {
		for i := 0; i < 50; i++ {
			time.Sleep(10000)
			a, b = b, Add(a, b)
			ret <- a
		}
	}()
	//for val := range ret {
	//	fmt.Println(val)
	//}
	listBuf := make([]int, 10)
	for {
		if len(listBuf) == 10 {
			fmt.Println(listBuf)
			listBuf = listBuf[:0]
		}
		val := <-ret
		listBuf = append(listBuf, val)
	}
}
