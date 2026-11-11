package metrics

import (
	"os"
	"testing"
)

// TestGetMetricsConfigURLDefaultEmpty proves KilatMesh Agent builds carry
// no built-in metrics config endpoint: without an explicit operator
// override there is no remote configuration to fetch.
func TestGetMetricsConfigURLDefaultEmpty(t *testing.T) {
	t.Setenv(EnvMetricsConfigURL, "")
	if err := os.Unsetenv(EnvMetricsConfigURL); err != nil {
		t.Fatalf("unset %s: %v", EnvMetricsConfigURL, err)
	}
	if got := getMetricsConfigURL(); got != "" {
		t.Errorf("default metrics config URL must be empty, got %q", got)
	}

	t.Setenv(EnvMetricsConfigURL, "https://metrics.example.internal/config")
	if got := getMetricsConfigURL(); got != "https://metrics.example.internal/config" {
		t.Errorf("env override must win, got %q", got)
	}
}
