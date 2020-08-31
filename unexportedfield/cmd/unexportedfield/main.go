package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/moricho/staticanalysis/unexportedfield"
)

func main() { unitchecker.Main(unexportedfield.Analyzer) }
