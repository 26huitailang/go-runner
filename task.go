package main

import (
	"fmt"
	"time"
)

type ITask interface {
	Run(c chan<- string)
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

func (t *Task) Run(c chan<- string) {
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%v closed", t.Name)
			c <- msg
		}
	}()
	t.Status = true
	for i := 0; ; i++ {
		c <- t.Name
		time.Sleep(1 * time.Second)
		if i > 5 {
			panic("zzzzzzz")
		}
		//fmt.Println(t.Name)
	}
}
