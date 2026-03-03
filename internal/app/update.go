package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/utils"
)

func (a *App) UpdateCmd() *cobra.Command {
	const usage = "update /path/to/data/files"
	update := &cobra.Command{
		Use:   usage,
		Short: "Update repo data",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				slog.Error(
					"invalid arguments",
					"usage", "updater ${data-file-dir}",
				)
				return fmt.Errorf("invalid argument")
			}
			return update(a.Client, a.Ctx, args[0])
		},
	}
	return update
}

func update(client *api.Client, ctx context.Context, dir string) error {
	list, err := io.GetDataFiles(dir)
	if err != nil {
		slog.Error(
			"failed to get data files",
			"dir", dir,
			"err", err,
		)
		return fmt.Errorf("failed to get data files at %s: %w", dir, err)
	}

	for _, path := range list {
		slog.Info(
			"processing",
			"path", path,
		)

		if err := updateSingleFile(client, ctx, path); err != nil {
			slog.Error(
				"update file failed",
				"path", path,
				"err", err,
			)
			return fmt.Errorf("update file %s failed: %w", path, err)
		}
	}

	slog.Info("finished")
	return nil
}

func updateSingleFile(client *api.Client, ctx context.Context, path string) error {
	const maxTagNumber = 5
	owner := utils.ParseOwner(path)

	projects, err := io.ReadYaml(path)
	if err != nil {
		return fmt.Errorf("read yaml %s: %w", path, err)
	}

	for i := range projects {
		slog.Info(
			"processing",
			"repo", owner+"/"+projects[i].Repo,
		)

		repo, status, err := client.GetRepo(ctx, owner, projects[i].Repo)
		if err != nil {
			return fmt.Errorf("get repo %s/%s: %w", owner, projects[i].Repo, err)
		}

		if status != http.StatusOK {
			return fmt.Errorf("get repo %s/%s failed: unexpected status %d", owner, projects[i].Repo, status)
		}

		tags := append([]string{repo.Language}, repo.Tags...)
		slog.Debug(
			"get tags",
			"size", len(tags),
		)

		if len(tags) > maxTagNumber {
			slog.Debug("remove some tags")
			tags = tags[:maxTagNumber]
		}
		projects[i].Description = repo.Description
		projects[i].Stars = repo.Stars
		projects[i].UpdatedAt = repo.UpdatedAt
		projects[i].Tags = tags
	}

	if err := io.WriteYaml(projects, path); err != nil {
		return fmt.Errorf("write yaml to %s: %w", path, err)
	}

	return nil
}
