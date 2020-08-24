package dupimport_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/moricho/staticanalysis/dupimport"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimport.Analyzer, "a")
}
