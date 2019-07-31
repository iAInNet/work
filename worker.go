package work

type Worker struct {
	Call           WorkerFunc
	MaxConcurrency int
}

type WorkerFunc interface {
	Run(task Task, args ...interface{}) TaskResult
}

type MyWorkerFunc func(task Task, args ...interface{}) (TaskResult)

func (f MyWorkerFunc) Run(task Task, args ...interface{}) (TaskResult) {
	return f(task, args)
}
