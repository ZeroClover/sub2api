package openai

import "testing"

func TestIsProOnlyModel(t *testing.T) {
	tests := []struct {
		name  string
		model string
		want  bool
	}{
		{name: "exact pro-only model", model: "gpt-5.3-codex-spark", want: true},
		{name: "pro-only model suffix variant", model: "gpt-5.3-codex-spark-high", want: true},
		{name: "pro-only model with path prefix", model: "openai/gpt-5.3-codex-spark", want: true},
		{name: "codex spark alias", model: "codex-spark", want: true},
		{name: "codex spark alias suffix", model: "codex-spark-xhigh", want: true},
		{name: "regular model", model: "gpt-5.3-codex", want: false},
		{name: "empty model", model: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProOnlyModel(tt.model); got != tt.want {
				t.Fatalf("IsProOnlyModel(%q) = %v, want %v", tt.model, got, tt.want)
			}
		})
	}
}
