package io

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/universtar-org/ust/internal/model"
)

func WriteYaml(projects []model.Project, path string) error {
	slog.Debug("write yaml start", "path", path, "count", len(projects))

	data, err := yaml.Marshal(projects)
	if err != nil {
		return fmt.Errorf("marshal yaml %s: %w", path, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write yaml %s: %w", path, err)
	}
	return nil
}
