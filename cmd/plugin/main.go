// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-python/internal/plugin"
)

func main() {
	publisher := plugin.NewPublisher(plugin.Config{})
	log.Printf("updater-python plugin ready: updates Python package metadata and uploads distributions (%T)", publisher)
}
