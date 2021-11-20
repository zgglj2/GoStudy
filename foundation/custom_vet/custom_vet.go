package main

import (
	"github.com/blanchonvincent/ctxarg/analysis/passes/ctxarg"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
)

func main() {
	multichecker.Main(
		atomic.Analyzer,
		loopclosure.Analyzer,
		ctxarg.Analyzer,
	)
}
