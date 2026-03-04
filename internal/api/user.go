package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/universtar-org/ust/internal/model"
)

func (c *Client) GetUser(ctx context.Context, username string) (*model.User, error) {
	slog.Debug("get user", "username", username)

	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/users/%s", username))
	if err != nil {
		return nil, fmt.Errorf("get user %s: %w", username, err)
	}

	var r model.User
	status, err := c.do(req, &r)
	if err != nil || status != http.StatusOK {
		if err != nil {
			return nil, fmt.Errorf("get user %s: %w", username, err)
		} else {
			return nil, fmt.Errorf("get user %s: unexpected status %d", username, status)
		}
	}

	return &r, nil
}
