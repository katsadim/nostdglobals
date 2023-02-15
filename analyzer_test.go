package nostdglobals_test

import (
	"testing"

	"github.com/katsadim/nostdglobals"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nostdglobals.Analyzer)
}
