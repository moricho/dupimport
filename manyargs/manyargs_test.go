package manyargs_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/moricho/staticanalysis/manyargs"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, manyargs.Analyzer, "a")
}
