package plugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/Nishiramirai/loglint/analyzer"
	"github.com/golangci/plugin-module-register/register"
)

func init() {
	register.Plugin("loglint", New)
}

type Settings struct{}

type Plugin struct{}

var _ register.LinterPlugin = (*Plugin)(nil)

func New(conf any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
