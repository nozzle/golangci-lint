package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/nozzle/golangci-lint/pkg/golinters/nozzle/extractorlint/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.HandlerAnalyzer)
}
