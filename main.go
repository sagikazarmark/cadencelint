package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/sagikazarmark/cadencelint/analysis/passes/determinism"
)

func main() {
	multichecker.Main(determinism.Analyzer)
}
