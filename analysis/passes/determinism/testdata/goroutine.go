package testdata

import "go.uber.org/cadence/workflow"

func GoRoutine(ctx workflow.Context) {
	go func() { // want "workflows must not start goroutines"

	}()
}
