package unexportedfield_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/moricho/staticanalysis/unexportedfield"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unexportedfield.Analyzer, "a")
}
