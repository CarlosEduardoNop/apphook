package workerpool

import (
	"fmt"
	"sync"
)

type Task interface {
	Process()
}

type CallTask struct {
	ID int
}

func (t *CallTask) Process() {
	fmt.Println("Processing task", t.ID)
}

type WorkerPool struct {
	Tasks      []Task
	concurency int
	tasksChan  chan Task
	wg         sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	for i:= 0; i < wp.concurency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	wp.wg.Wait()
}
