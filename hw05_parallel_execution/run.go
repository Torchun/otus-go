package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// n = goroutines count
	// m = max errors

	// prepare & run n goroutines with single task per each

	var wg sync.WaitGroup
	for index, task := range tasks {
		fmt.Println(index, task)
		// fill queue with tasks
		// goroutines will take and process each task
		// seems like channel should be blocked until all tasks are taken
	}

	wg.Wait()
	return nil
}

func Runner() {
	// do single task
}
