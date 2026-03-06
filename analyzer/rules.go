package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func checkLowercase(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkLowercase || msg == "" {
		return
	}

	first := rune(msg[0])
	if unicode.IsUpper(first) {
		pass.Reportf(node.Pos(), "log message should start with lowercase letter")
	}
}

func checkEnglish(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkEnglish {
		return
	}

	for _, r := range msg {
		if r > 127 {
			pass.Reportf(node.Pos(), "log message should contain only English characters")
			return
		}
	}
}

var specialCharRegex = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?~]|[\p{So}]`)

func checkSpecialChars(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkSpecialChars {
		return
	}

	if specialCharRegex.MatchString(msg) {
		pass.Reportf(node.Pos(), "log message should not contain special characters or emojis")
	}
}

func checkSensitive(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkSensitive {
		return
	}

	lower := strings.ToLower(msg)
	for _, keyword := range cfg.sensitiveKeywords {
		if strings.Contains(lower, strings.ToLower(keyword)) {
			pass.Reportf(node.Pos(), "log message contains sensitive data")
			return
		}
	}
}