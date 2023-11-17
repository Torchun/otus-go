package hw06pipelineexecution

import (
	"fmt"
	"reflect"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	// each call got current (!same!) value
	recv := in
	fmt.Println("recv: ", reflect.TypeOf(recv))
	fmt.Println(recv)

	for _, stage := range stages {

		stg := stage(recv)
		fmt.Println("stg: ", reflect.TypeOf(stg))
		fmt.Println(stg)

		recv = executeStage(stg, done)
		fmt.Println("stage: ", reflect.TypeOf(stage))
		fmt.Println(stage)

	}
	return recv
}

func executeStage(recv In, done In) Out {
	send := make(Bi)
	go func() {
		defer close(send)
		fmt.Println("func send: ", reflect.TypeOf(send))
		fmt.Println(send)
		for {
			select {
			case v, ok := <-recv:
				fmt.Println("v, ok: ", reflect.TypeOf(v), ok)
				fmt.Println(v)
				if !ok {
					return
				}
				fmt.Println(ok, v)
				send <- v
			case <-done:
				fmt.Println("done: ", reflect.TypeOf(done))
				fmt.Println(done)
				return
			}
		}
	}()
	return send
}
