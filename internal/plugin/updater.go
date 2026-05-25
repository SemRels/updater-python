// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

// Package plugin updates pyproject.toml files in-place.
package plugin

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var versionPattern = regexp.MustCompile(`^(\s*version\s*=\s*)"[^"]*"(\s*)$`)

// Updater updates Python project versions.
type Updater struct{}

// NewUpdater creates an updater.
func NewUpdater() *Updater {
	return &Updater{}
}

// Update rewrites the version line in [project] or [tool.poetry].
func (u *Updater) Update(path, version string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	updated, err := updateContent(string(data), version)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func updateContent(content, version string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	lines := make([]string, 0)
	inSection := false
	updated := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "[") {
			lower := strings.ToLower(trimmed)
			inSection = lower == "[project]" || lower == "[tool.poetry]"
		}

		if inSection && !updated && versionPattern.MatchString(line) {
			line = versionPattern.ReplaceAllString(line, `${1}"`+version+`"${2}`)
			updated = true
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scan pyproject.toml: %w", err)
	}
	if !updated {
		return "", fmt.Errorf("version field not found in pyproject.toml")
	}
	return strings.Join(lines, "\n"), nil
}
