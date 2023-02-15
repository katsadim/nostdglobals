package main

import (
	"github.com/katsadim/nostdglobals"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nostdglobals.Analyzer)
}
