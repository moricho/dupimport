package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/moricho/staticanalysis/dupimport"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
