package main

func run(tasks []ITask, irunner IRunner) {
	irunner.Register(tasks)
	irunner.Run()
}

func main() {
	// todo: 是不是可以每个任务返回一个channel，主程序去select这些返回状态
	t1 := NewTask("t1", false)
	//t2 := NewTask("t2", false)
	t2 := NewTaskAdd()
	tasks := []ITask{t1, t2}
	c := make(chan string)
	runner := &Runner{chStatus: c}
	run(tasks, runner)
}
