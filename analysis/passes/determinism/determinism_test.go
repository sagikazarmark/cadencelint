package determinism

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDeterminism(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
