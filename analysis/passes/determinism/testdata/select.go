package testdata

import "go.uber.org/cadence/workflow"

var ch chan struct{}

func Select(ctx workflow.Context) {
	select { // want "workflows must not select directly from channels"
		case <-ch:
	}
}
