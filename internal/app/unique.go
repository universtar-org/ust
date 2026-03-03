package app

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/model"
)

func (a *App) UniqueCmd() *cobra.Command {
	unique := &cobra.Command{
		Use:   "unique [username]",
		Short: "Check uniqueness of a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("username required")
			}
			if err := unique(a.Client, a.Ctx, args[0]); err != nil {
				return fmt.Errorf("check unique of %s failed: %w", args[0], err)
			}
			return nil
		},
	}

	return unique
}

func unique(client *api.Client, ctx context.Context, username string) error {
	repoOwner := "universtar-org"
	slog.Info(
		"check uniquess",
		"user", username,
	)

	user, err := checkUsername(client, ctx, username)
	if err != nil {
		slog.Error(
			"check username failed",
			"err", err,
		)
		return fmt.Errorf("check user %s failed: %w", username, err)
	}

	repos, err := client.GetRepoByUser(ctx, repoOwner)
	if err != nil {
		slog.Error(
			"get repo by user failed",
			"user", username,
			"err", err,
		)
		return fmt.Errorf("get repo of %s failed: %w", username, err)
	}

	if err := checkUniqueness(client, ctx, repos, *user, repoOwner); err != nil {
		slog.Error(
			"check uniqueness",
			"user", username,
			"err", err,
		)
		return fmt.Errorf("check uniqueness of %s failed: %w", username, err)
	}
	slog.Info("finished")

	return nil
}

func checkUsername(client *api.Client, ctx context.Context, username string) (*model.User, error) {
	user, err := client.GetUser(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user %s: %w", username, err)
	}
	if username != user.Name {
		return nil, fmt.Errorf("username mismatch: get %s, expect: %s", username, user.Name)
	}
	return user, nil
}

func checkUniqueness(client *api.Client, ctx context.Context, repos []model.Repo, user model.User, repoOwner string) error {
	projectWhiteList := []string{"tools", "www"}

	if user.Type != "User" {
		return nil
	}

	path := "data/projects"
	for _, repo := range repos {
		if slices.Contains(projectWhiteList, repo.Name) {
			continue
		}

		slog.Info(
			"checking",
			"repo", repoOwner+"/"+repo.Name,
		)
		contents, err := client.GetDirContent(ctx, repoOwner, repo.Name, path)
		if err != nil {
			return fmt.Errorf("get dir content %s/%s/%s: %w", repoOwner, repo.Name, path, err)
		}

		for _, content := range contents {
			if user.Name == strings.TrimSuffix(content, filepath.Ext(content)) {

				return fmt.Errorf("duplicated username in %s/%s", repoOwner, repo.Name)
			}
		}
	}

	return nil
}
