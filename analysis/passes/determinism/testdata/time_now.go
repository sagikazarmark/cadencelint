package testdata

import (
	"time"

	"go.uber.org/cadence/workflow"
)

func TimeNow(ctx workflow.Context) {
	_ = time.Now() // want "workflows must not call time.Now and time.Sleep"
}
