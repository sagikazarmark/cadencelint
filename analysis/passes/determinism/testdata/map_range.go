package testdata

import "go.uber.org/cadence/workflow"

func WorkflowWithMapRange(ctx workflow.Context) {
	myMap := map[string]string{
		"key": "value",
	}

	for range myMap { // want "workflows must not range over maps"
		// noop
	}
}
