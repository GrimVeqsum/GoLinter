package analyzer

import (
	"flag"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	cfgPath := ""
	a := &analysis.Analyzer{
		Name: "loglint",
		Doc:  "checks log messages for style and sensitive data rules",
		Run: func(pass *analysis.Pass) (any, error) {
			cfg, err := loadConfig(cfgPath)
			if err != nil {
				return nil, err
			}
			return run(pass, cfg), nil
		},
	}
	a.Flags = *flag.NewFlagSet(a.Name, flag.ExitOnError)
	a.Flags.StringVar(&cfgPath, "config", "", "path to loglint JSON config file")
	return a
}

var Analyzer = NewAnalyzer()

func run(pass *analysis.Pass, cfg runtimeConfig) any {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			ident, ok := sel.X.(*ast.Ident)
			if !ok {
				return true
			}

			if ident.Name != "log" && ident.Name != "slog" && ident.Name != "zap" {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			msgArg, ok := call.Args[0].(*ast.BasicLit)
			if !ok || msgArg.Kind != token.STRING {
				return true
			}

			msg := strings.Trim(msgArg.Value, "\"")

			checkLowercase(pass, msgArg, msg, cfg)
			checkEnglish(pass, msgArg, msg, cfg)
			checkSpecialChars(pass, msgArg, msg, cfg)
			checkSensitive(pass, msgArg, msg, cfg)

			return true
		})
	}
	return nil
}