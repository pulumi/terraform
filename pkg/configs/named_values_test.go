package configs

import (
	"testing"

	"github.com/spf13/afero"
)

func expectTrue(t *testing.T, condition bool, message string) {
	t.Helper()
	if !condition {
		t.Fatalf("expected true: %s", message)
	}
}

func TestParsingEphemeralVariables(t *testing.T) {
	t.Parallel()

	src := `
variable "foo_wo" {
	type = string
	sensitive = true
	ephemeral = true
}`

	fs := afero.NewMemMapFs()
	if err := afero.WriteFile(fs, "variables.tf", []byte(src), 0600); err != nil {
		t.Fatalf("unable to write test file: %s", err)
	}

	parser := NewParser(fs)
	module, diags := parser.LoadConfigDir(".")
	if diags.HasErrors() {
		t.Errorf("unexpected error diagnostics")
		for _, diag := range diags {
			t.Logf("- %s", diag)
		}
	}

	expectTrue(t, len(module.Variables) == 1, "expected exactly one variable")
	variable, exists := module.Variables["foo_wo"]
	expectTrue(t, exists, "variable foo_wo should be present")
	expectTrue(t, variable.EphemeralSet, "variable foo_wo EphemeralSet should be true")
	expectTrue(t, variable.Ephemeral, "variable foo_wo Ephemeral should be true")
}
