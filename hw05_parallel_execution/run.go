package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) isGreaterOrEqual(value int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value >= value
}

func Run(tasks []Task, tasksThreads, errorsLimit int) error {
	if tasksThreads < 1 {
		return nil
	}
	if errorsLimit <= 0 {
		errorsLimit = len(tasks)
	}
	errorsCount := Counter{value: 0}
	tasksChannel := make(chan Task)
	wg := sync.WaitGroup{}

	for i := 0; i < tasksThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				err := task()
				if err != nil {
					errorsCount.inc()
				}
			}
		}()
	}

	for _, task := range tasks {
		if errorsCount.isGreaterOrEqual(errorsLimit) {
			break
		}
		tasksChannel <- task
	}
	close(tasksChannel)

	wg.Wait()

	if errorsCount.isGreaterOrEqual(errorsLimit) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
