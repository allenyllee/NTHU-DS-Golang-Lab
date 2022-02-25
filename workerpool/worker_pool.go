package workerpool

import (
	"context"
	"log"
	"os"
	"sync"
)

type Task struct {
	Func func(args ...interface{}) *Result
	Args []interface{}
}

type Result struct {
	Value interface{}
	Err   error
}

type WorkerPool interface {
	Start(ctx context.Context)
	Tasks() chan *Task
	Results() chan *Result
}

type workerPool struct {
	numWorkers int
	tasks      chan *Task
	results    chan *Result
	wg         *sync.WaitGroup
}

var _ WorkerPool = (*workerPool)(nil)
var logger *log.Logger
type ctxKey struct{}
var key ctxKey = ctxKey{}

func NewWorkerPool(numWorkers int, bufferSize int) *workerPool {
	return &workerPool{
		numWorkers: numWorkers,
		tasks:      make(chan *Task, bufferSize),
		results:    make(chan *Result, bufferSize),
		wg:         &sync.WaitGroup{},
	}
}

func (wp *workerPool) Start(ctx context.Context) {
	// TODO: implementation
	//
	// Starts numWorkers of goroutines, wait until all jobs are done.
	// Remember to closed the result channel before exit.
	logger = log.New(os.Stdout, "", log.Ltime)

	logger.Println("任務", ctx.Value("name"), ":任務========")

	for i := 0; i < wp.numWorkers; i++ {
		valueCtx := context.WithValue(ctx, key, i)
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			wp.run(valueCtx)
		}()
	}

	wp.wg.Wait()
	close(wp.Results())
}

func (wp *workerPool) Tasks() chan *Task {
	return wp.tasks
}

func (wp *workerPool) Results() chan *Result {
	return wp.results
}

func (wp *workerPool) run(ctx context.Context) {
	// TODO: implementation
	//
	// Keeps fetching task from the task channel, do the task,
	// then makes sure to exit if context is done.
	for {
		select {
		case <- ctx.Done():
			logger.Println("任務", ctx.Value(key), ":任務取消...")
			return
		default:
		}

		select {
		case <- ctx.Done():
			logger.Println("任務", ctx.Value(key), ":任務取消...")
			return
		case task, ok := <-wp.Tasks():
			if ok {
				logger.Println("任務", ctx.Value(key), ":工作中")
				wp.Results() <- task.Func(task.Args...)
			} else {
				logger.Println("任務", ctx.Value(key), ":無任務")
				return
			}
		}
	}
}
