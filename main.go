package main

func run(tasks []*Task, irunner IRunner) {
	irunner.Register(tasks)
	irunner.Run()
}

func main() {
	t1 := NewTask("t1", false)
	t2 := NewTask("t2", false)
	tasks := []*Task{t1, t2}
	c := make(chan string)
	runner := &Runner{chStatus: c}
	run(tasks, runner)
}
