package testdata

import (
	"time"

	"go.uber.org/cadence/workflow"
)

func TimeSleep(ctx workflow.Context) {
	time.Sleep(time.Second) // want "workflows must not call time.Now and time.Sleep"
}
