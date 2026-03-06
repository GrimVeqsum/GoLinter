package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer_DefaultRules(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, NewAnalyzer(), "logrules")
}

func TestAnalyzer_WithConfigDisablesRules(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer()
	if err := a.Flags.Set("config", filepath.Join(testdata, "loglint.disabled.json")); err != nil {
		t.Fatalf("set config flag: %v", err)
	}
	analysistest.Run(t, testdata, a, "logrulesdisabled")
}

func TestAnalyzer_WithMissingConfigFile_UsesDefaults(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer()
	if err := a.Flags.Set("config", filepath.Join(testdata, "does-not-exist.json")); err != nil {
		t.Fatalf("set config flag: %v", err)
	}
	analysistest.Run(t, testdata, a, "logrules")
}