package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	// n = goroutines count
	// m = max errors

	// prepare & run n goroutines with single task per each

	var wg sync.WaitGroup
	// all channels are sync, writer/reader wil be waiting for it's counterpart
	tasksChannel := make(chan Task)
	errorsChannel := make(chan struct{}, m)
	exitChannel := make(chan bool, n)

	// don't forget to close channels
	defer close(tasksChannel)
	defer close(errorsChannel)
	defer close(exitChannel)

	// run each task in separate runner in parallel
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Runner(&wg, m, tasksChannel, errorsChannel, exitChannel) // wg should be same for all goroutines
	}

	for index, task := range tasks {
		// write to 0-len queue
		// goroutines will take and process each task one by one
		// channel would be blocked until all tasks are taken
		if n+m == index {
			if len(errorsChannel) == m {
				// exit from loop if err cap exceeded m
				err = ErrErrorsLimitExceeded
				break
			}
		}
		fmt.Println("Run: ", index, &task)
		// will be waiting until someone take proposed task, sequentially
		tasksChannel <- task
		fmt.Println("tasksChannel: ", len(tasksChannel))
	}

	// means all tasks are taken
	// executing runner will be switched to read from channel
	// for those who is not - need to send exit signal
	for i := 0; i < n; i++ {
		exitChannel <- true
	}
	wg.Wait()

	// check if m cap was exceeded
	if err != nil {
		return err
	}

	return nil
}

func Runner(wg *sync.WaitGroup, m int, tasksChannel chan Task, errorsChannel chan struct{}, exitChannel chan bool) {
	// do single task
	// don't forget to close self in WaitGroup
	defer wg.Done()
	fmt.Println("tasksChannel: ", len(tasksChannel))

	// infinite loop to read from channel
	for {
		// round robin on performable cases
		select {
		// if task can be read
		case task, ok := <-tasksChannel:
			// check if channel is closed - desirable check
			if !ok {
				return
			}
			// execute task (task is a function)
			err := task()
			fmt.Println("runner: ", m, ok, err)
			// check if error retrieved
			if err != nil {
				// exit if error count exceeded allowed cap of m errors
				if len(errorsChannel) == m {
					return
				}
				// otherwise add +1 error to channel
				errorsChannel <- struct{}{}
			}
		// if can be read - exit the goroutine
		case <-exitChannel:
			return
		}
	}
}
