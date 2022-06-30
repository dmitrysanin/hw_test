package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type taskErrCounter struct {
	mu       sync.Mutex
	errCount int
}

func (tec *taskErrCounter) getErrCount() int {
	tec.mu.Lock()
	defer tec.mu.Unlock()
	return tec.errCount
}

func (tec *taskErrCounter) incrementErrCount() {
	tec.mu.Lock()
	defer tec.mu.Unlock()
	tec.errCount++
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tec := taskErrCounter{}

	wg := sync.WaitGroup{}
	wg.Add(n)
	taskChanel := make(chan Task)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for {
				t, ok := <-taskChanel
				if ok {
					if errCount := tec.getErrCount(); errCount < m {
						err := t()
						if err != nil {
							tec.incrementErrCount()
						}
					}
				} else {
					return
				}
			}
		}()
	}

	for _, t := range tasks {
		taskChanel <- t
	}
	close(taskChanel)

	wg.Wait()

	if errCount := tec.getErrCount(); errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
