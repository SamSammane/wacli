package main

import (
	"strings"
	"testing"
)

func TestRootFlagsReadOnlyFlag(t *testing.T) {
	flags := &rootFlags{readOnly: true}

	if !flags.isReadOnly() {
		t.Fatal("isReadOnly = false, want true")
	}
	err := flags.requireWritable()
	if err == nil || !strings.Contains(err.Error(), "read-only mode") {
		t.Fatalf("requireWritable error = %v", err)
	}
}

func TestRootFlagsReadOnlyEnv(t *testing.T) {
	t.Setenv("WACLI_READONLY", "yes")

	if !(&rootFlags{}).isReadOnly() {
		t.Fatal("isReadOnly = false, want true")
	}
}
