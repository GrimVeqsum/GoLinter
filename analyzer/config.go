package analyzer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var defaultSensitiveKeywords = []string{"password", "token", "api_key", "secret"}

type fileConfig struct {
	DisableRules      []string `json:"disable_rules"`
	SensitiveKeywords []string `json:"sensitive_keywords"`
}

type runtimeConfig struct {
	checkLowercase    bool
	checkEnglish      bool
	checkSpecialChars bool
	checkSensitive    bool
	sensitiveKeywords []string
}

func defaultConfig() runtimeConfig {
	return runtimeConfig{
		checkLowercase:    true,
		checkEnglish:      true,
		checkSpecialChars: true,
		checkSensitive:    true,
		sensitiveKeywords: append([]string(nil), defaultSensitiveKeywords...),
	}
}

func loadConfig(path string) (runtimeConfig, error) {
	cfg := defaultConfig()
	if path == "" {
		return cfg, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cfg, nil
		}
		return runtimeConfig{}, fmt.Errorf("read config file: %w", err)
	}

	var parsed fileConfig
	if err := json.Unmarshal(content, &parsed); err != nil {
		return runtimeConfig{}, fmt.Errorf("parse config file: %w", err)
	}

	disable := make(map[string]struct{}, len(parsed.DisableRules))
	for _, rule := range parsed.DisableRules {
		disable[strings.ToLower(strings.TrimSpace(rule))] = struct{}{}
	}

	cfg.checkLowercase = !contains(disable, "lowercase")
	cfg.checkEnglish = !contains(disable, "english")
	cfg.checkSpecialChars = !contains(disable, "special_chars")
	cfg.checkSensitive = !contains(disable, "sensitive")

	if len(parsed.SensitiveKeywords) > 0 {
		cfg.sensitiveKeywords = parsed.SensitiveKeywords
	}

	return cfg, nil
}

func contains(set map[string]struct{}, key string) bool {
	_, ok := set[key]
	return ok
}