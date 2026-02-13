// Package openai provides helpers and types for OpenAI API integration.
package openai

import (
	_ "embed"
	"strings"
)

// Model represents an OpenAI model
type Model struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Created     int64  `json:"created"`
	OwnedBy     string `json:"owned_by"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
}

// DefaultModels OpenAI models list
var DefaultModels = []Model{
	{ID: "gpt-5.3", Object: "model", Created: 1735689600, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.3"},
	{ID: "gpt-5.3-codex-spark", Object: "model", Created: 1739404800, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.3 Codex Spark"},
	{ID: "gpt-5.3-codex", Object: "model", Created: 1735689600, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.3 Codex"},
	{ID: "gpt-5.2", Object: "model", Created: 1733875200, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.2"},
	{ID: "gpt-5.2-codex", Object: "model", Created: 1733011200, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.2 Codex"},
	{ID: "gpt-5.1-codex-max", Object: "model", Created: 1730419200, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.1 Codex Max"},
	{ID: "gpt-5.1-codex", Object: "model", Created: 1730419200, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.1 Codex"},
	{ID: "gpt-5.1", Object: "model", Created: 1731456000, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.1"},
	{ID: "gpt-5.1-codex-mini", Object: "model", Created: 1730419200, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5.1 Codex Mini"},
	{ID: "gpt-5", Object: "model", Created: 1722988800, OwnedBy: "openai", Type: "model", DisplayName: "GPT-5"},
}

// DefaultModelIDs returns the default model ID list
func DefaultModelIDs() []string {
	ids := make([]string, len(DefaultModels))
	for i, m := range DefaultModels {
		ids[i] = m.ID
	}
	return ids
}

// DefaultTestModel default model for testing OpenAI accounts
const DefaultTestModel = "gpt-5.1-codex"

// ProOnlyModels 仅 ChatGPT Pro 订阅可用的模型
var ProOnlyModels = map[string]bool{
	"gpt-5.3-codex-spark": true,
}

// IsProOnlyModel 检查模型是否仅限 Pro 订阅
func IsProOnlyModel(model string) bool {
	normalized := normalizeProOnlyModel(model)
	if normalized == "" {
		return false
	}
	if ProOnlyModels[normalized] {
		return true
	}
	for proModel := range ProOnlyModels {
		if strings.HasPrefix(normalized, proModel+"-") {
			return true
		}
	}
	// Codex CLI 场景常见别名（如 codex-spark / codex-spark-high）。
	return normalized == "codex-spark" || strings.HasPrefix(normalized, "codex-spark-")
}

func normalizeProOnlyModel(model string) string {
	normalized := strings.TrimSpace(strings.ToLower(model))
	if normalized == "" {
		return ""
	}
	if strings.Contains(normalized, "/") {
		parts := strings.Split(normalized, "/")
		normalized = parts[len(parts)-1]
	}
	normalized = strings.ReplaceAll(normalized, " ", "-")
	return normalized
}

// DefaultInstructions default instructions for non-Codex CLI requests
// Content loaded from instructions.txt at compile time
//
//go:embed instructions.txt
var DefaultInstructions string
