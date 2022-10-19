package lib

// 具体任务,可以传参可以自定义操作
type PoolTask struct {
	Args interface{}
	Do   func(interface{}) error
}

// 任务通道
var PoolJobChannels = make(chan PoolTask)

// 入口的任务通道
var PoolJobs = make(chan PoolTask)

// 执行
func PoolRun(num int) {
	for i := 0; i < num; i++ {
		go worker(i)
	}
	for task := range PoolJobs {
		PoolJobChannels <- task
	}
	close(PoolJobChannels)
}

// 实际的工作协程worker
func worker(workId int) {
	for task := range PoolJobChannels {
		task.Do(task.Args)
	}
}
