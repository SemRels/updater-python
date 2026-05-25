package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdaterUpdateProjectVersion(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-python-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "pyproject.toml")
	original := "[project]\nname = \"demo\"\nversion = \"1.2.3\"\n"
	if err := os.WriteFile(file, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "1.3.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `version = "1.3.0"`) {
		t.Fatalf("updated file = %s", got)
	}
}

func TestUpdaterUpdatePoetryVersion(t *testing.T) {
	t.Parallel()

	updated, err := updateContent("[tool.poetry]\nversion = \"1.0.0\"\n", "1.1.0")
	if err != nil {
		t.Fatalf("updateContent() error = %v", err)
	}
	if !strings.Contains(updated, `version = "1.1.0"`) {
		t.Fatalf("updated content = %s", updated)
	}
}

func TestUpdaterMissingFile(t *testing.T) {
	t.Parallel()

	err := NewUpdater().Update(filepath.Join(t.TempDir(), "pyproject.toml"), "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "read") {
		t.Fatalf("expected read error, got %v", err)
	}
}

func TestUpdaterMissingVersionField(t *testing.T) {
	t.Parallel()

	_, err := updateContent("[project]\nname = \"demo\"\n", "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "version field not found") {
		t.Fatalf("expected version error, got %v", err)
	}
}
