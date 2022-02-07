package main

import (
	"github.com/katsadim/nostdglobals/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
