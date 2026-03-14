package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "log linter",
	Run:  run,
}
