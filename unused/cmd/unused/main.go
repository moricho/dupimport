package main

import (
	"github.com/moricho/unused"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unused.Analyzer) }
