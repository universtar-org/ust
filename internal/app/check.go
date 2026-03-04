package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/universtar-org/ust/internal/api"
	"github.com/universtar-org/ust/internal/io"
	"github.com/universtar-org/ust/internal/utils"
)

func (a *App) CheckCmd() *cobra.Command {
	const usage = "check /path/to/data/files"
	check := &cobra.Command{
		Use:   usage,
		Short: "Check repo files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				slog.Error(
					"invalid argument",
					"usage", usage,
				)
				return fmt.Errorf("invalid argument")
			}
			return check(a.Client, a.Ctx, args[0])
		},
	}

	return check
}

func check(client *api.Client, ctx context.Context, dir string) error {
	list, err := io.GetDataFiles(dir)
	if err != nil {
		slog.Error(
			"failed to get data files",
			"dir", dir,
			"err", err,
		)
		return fmt.Errorf("failed to read dir %s: %w", dir, err)
	}

	for _, path := range list {
		slog.Info(
			"checking file",
			"path", path,
		)
		if err := checkSingleFile(client, ctx, path); err != nil {
			slog.Error(
				"check failed",
				"path", path,
				"err", err,
			)
			return fmt.Errorf("check file %s failed: %w", path, err)
		}
	}

	slog.Info("finished")
	return nil
}

func checkSingleFile(client *api.Client, ctx context.Context, path string) error {
	projects, err := io.ReadYaml(path)
	if err != nil {
		return fmt.Errorf("read yaml %s: %w", path, err)
	}

	owner := utils.ParseOwner(path)

	for _, project := range projects {
		slog.Debug("checking repo",
			"owner", owner,
			"repo", project.Repo,
		)

		_, status, err := client.GetRepo(ctx, owner, project.Repo)
		if err != nil {
			return fmt.Errorf("checking repo %s/%s: %w", owner, project.Repo, err)
		}

		if status != http.StatusOK {
			return fmt.Errorf("check repo %s/%s: unexpected status %d", owner, project.Repo, status)
		}
	}

	return nil
}
