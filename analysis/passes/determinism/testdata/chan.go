package testdata

import "go.uber.org/cadence/workflow"

func Chan(ctx workflow.Context) {
	var _ chan struct{} // want "workflows must not create channels directly"
}
