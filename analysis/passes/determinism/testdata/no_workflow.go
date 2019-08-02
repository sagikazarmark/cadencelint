package testdata

import (
	"context"
	"time"
)

func NoWorkflow(ctx context.Context) {
	var ch chan struct{}

	select {
	case <-ch:
	}

	go func() {

	}()

	myMap := map[string]string{
		"key": "value",
	}

	for range myMap {
		// noop
	}

	_ = time.Now()
	time.Sleep(time.Second)
}
