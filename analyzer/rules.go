package analyzer

import (
	"go/ast"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

func reportWithFix(pass *analysis.Pass, node *ast.BasicLit, message, fixedText string) {
	pass.Report(analysis.Diagnostic{
		Pos:     node.Pos(),
		End:     node.End(),
		Message: message,
		SuggestedFixes: []analysis.SuggestedFix{{
			Message: "autofix log message",
			TextEdits: []analysis.TextEdit{{
				Pos:     node.Pos(),
				End:     node.End(),
				NewText: []byte(strconv.Quote(fixedText)),
			}},
		}},
	})
}

func checkLowercase(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkLowercase || msg == "" {
		return
	}

	r, size := utf8.DecodeRuneInString(msg)
	if unicode.IsUpper(r) {
		fixed := string(unicode.ToLower(r)) + msg[size:]
		reportWithFix(pass, node, "log message should start with lowercase letter", fixed)
	}
}

func checkEnglish(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkEnglish {
		return
	}

	var b strings.Builder
	hasNonASCII := false
	for _, r := range msg {
		if r > 127 {
			hasNonASCII = true
			continue
		}
		b.WriteRune(r)
	}
	if hasNonASCII {
		reportWithFix(pass, node, "log message should contain only English characters", b.String())
	}
}

var specialCharRegex = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?~]|[\p{So}]`)

func checkSpecialChars(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkSpecialChars {
		return
	}

	if specialCharRegex.MatchString(msg) {
		fixed := specialCharRegex.ReplaceAllString(msg, "")
		reportWithFix(pass, node, "log message should not contain special characters or emojis", fixed)
	}
}

func checkSensitive(pass *analysis.Pass, node *ast.BasicLit, msg string, cfg runtimeConfig) {
	if !cfg.checkSensitive {
		return
	}

	fixed := msg
	found := false
	for _, keyword := range cfg.sensitiveKeywords {
		if keyword == "" {
			continue
		}
		pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(keyword))
		if pattern.MatchString(fixed) {
			found = true
			fixed = pattern.ReplaceAllString(fixed, "[redacted]")
		}
	}

	if found {
		reportWithFix(pass, node, "log message contains sensitive data", fixed)
	}
}