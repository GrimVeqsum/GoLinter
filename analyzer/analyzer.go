package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
    Name: "loglint",
    Doc:  "checks log messages for style and sensitive data rules",
    Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
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

            // Поддерживаем log/slog и zap
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

            checkLowercase(pass, msgArg, msg)
            checkEnglish(pass, msgArg, msg)
            checkSpecialChars(pass, msgArg, msg)
            checkSensitive(pass, msgArg, msg)

            return true
        })
    }
    return nil, nil
}