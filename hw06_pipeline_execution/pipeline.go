package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// each call got current (!same!) value
	recv := in

	for _, stage := range stages {
		// execute each stage computing
		stg := stage(recv)
		// recv will contain formatted result of prev.stage calculations
		// executeUnderControl should be able to stop on signal in done channel
		recv = executeUnderControl(stg, done)
	}
	return recv
}

func executeUnderControl(stg In, done In) Out {
	// to comply with expected type
	send := make(Bi)
	go func() {
		defer close(send)
		// infinite wait until (stg or done) chan can be read
		for {
			select {
			// if can read from stage (one int per stage)
			case v, ok := <-stg:
				if !ok {
					// after once read will end goroutine
					return
				}
				// will be done once
				send <- v
			// for "done case" in pipeline_test - return immediately when got signal
			// that's why executeUnderControl needed
			case <-done:
				return
			}
		}
	}()
	return send
}
