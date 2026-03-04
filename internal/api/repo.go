package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/universtar-org/ust/internal/model"
)

// GetRepo Get repo information including description, number of stars, etc., via GitHub API.
func (c *Client) GetRepo(ctx context.Context, owner, repo string) (*model.Repo, int, error) {
	slog.Debug("get repo", "repo", owner+"/"+repo)

	url := fmt.Sprintf("/repos/%s/%s", owner, repo)
	req, err := c.newRequest(ctx, http.MethodGet, url)
	if err != nil {
		return nil, 0, err
	}

	var r model.Repo
	status, err := c.do(req, &r)
	if err != nil {
		return nil, status, err
	}

	return &r, status, nil
}

func (c *Client) GetRepoByUser(ctx context.Context, username string) ([]model.Repo, error) {
	slog.Debug("get repo by user", "user", username)

	url := fmt.Sprintf("/users/%s/repos", username)
	req, err := c.newRequest(ctx, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("get repo by user %s: %w", username, err)
	}

	var repos []model.Repo
	status, err := c.do(req, &repos)
	if err != nil {
		return nil, fmt.Errorf("get repo by user %s: %w", username, err)
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf("get repo by user %s: unexpected status %d", username, status)
	}

	return repos, nil
}

func (c *Client) GetDirContent(ctx context.Context, username, repo, path string) ([]string, error) {
	slog.Debug("get dir content", "repo", username+"/"+repo, "path", path)

	url := fmt.Sprintf("/repos/%s/%s/contents/%s", username, repo, path)
	req, err := c.newRequest(ctx, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("get dir content %s/%s/%s: %w", username, repo, path, err)
	}

	var contents []map[string]any
	var result []string
	status, err := c.do(req, &contents)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("get dir content %s/%s/%s: unexpected status %d", username, repo, path, status)
	}

	for _, v := range contents {
		name, ok := v["name"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid content item: missing name")
		}
		result = append(result, name)
	}

	return result, nil
}
