package golangciplugin

import (
	"fmt"

	"github.com/GrimVeqsum/GoLinter/analyzer"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type plugin struct{}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.NewAnalyzer()}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func init() {
	register.Plugin("loglint", New)
}

func New(conf any) (register.LinterPlugin, error) {
	if conf != nil {
		if _, ok := conf.(map[string]any); !ok {
			return nil, fmt.Errorf("loglint plugin settings must be a map")
		}
	}
	return &plugin{}, nil
}
