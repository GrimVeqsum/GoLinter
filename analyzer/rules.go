package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// 1. Проверка на строчную букву
func checkLowercase(pass *analysis.Pass, node *ast.BasicLit, msg string) {
    if msg == "" {
        return
    }
    first := rune(msg[0])
    if unicode.IsUpper(first) {
        pass.Reportf(node.Pos(), "log message should start with lowercase letter")
    }
}

// 2. Проверка на английский язык
func checkEnglish(pass *analysis.Pass, node *ast.BasicLit, msg string) {
    for _, r := range msg {
        if r > 127 {
            pass.Reportf(node.Pos(), "log message should contain only English characters")
            return
        }
    }
}

// 3. Проверка на спецсимволы и эмодзи
var specialCharRegex = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?~]|[\p{So}]`)

func checkSpecialChars(pass *analysis.Pass, node *ast.BasicLit, msg string) {
    if specialCharRegex.MatchString(msg) {
        pass.Reportf(node.Pos(), "log message should not contain special characters or emojis")
    }
}

// 4. Проверка на чувствительные данные
var sensitiveKeywords = []string{"password", "token", "api_key", "secret"}

func checkSensitive(pass *analysis.Pass, node *ast.BasicLit, msg string) {
    lower := strings.ToLower(msg)
    for _, keyword := range sensitiveKeywords {
        if strings.Contains(lower, keyword) {
            pass.Reportf(node.Pos(), "log message contains sensitive data")
        }
    }
}