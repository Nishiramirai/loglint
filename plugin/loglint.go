package plugin

import (
	"golang.org/x/tools/go/analysis"
	"github.com/Nishiramirai/loglint/analyzer"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
