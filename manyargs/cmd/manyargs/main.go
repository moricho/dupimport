package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/moricho/staticanalysis/manyargs"
)

func main() { unitchecker.Main(manyargs.Analyzer) }
