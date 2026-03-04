package io

import (
	"github.com/goccy/go-yaml"
	"github.com/universtar-org/ust/internal/model"

	"log/slog"
	"os"
	"path/filepath"
)

func ReadYaml(path string) ([]model.Project, error) {
	slog.Debug("read yaml start", "path", path)

	var projects []model.Project
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	slog.Debug("read yaml done", "path", path, "count", len(projects))

	return projects, nil
}

func GetDataFiles(dir string) ([]string, error) {
	slog.Debug("get all data file start", "dir", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, entry := range entries {
		if !entry.IsDir() {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}

	slog.Debug("get data file done", "dir", dir, "count", len(paths))

	return paths, nil
}
