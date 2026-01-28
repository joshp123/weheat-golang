package weheat

import "context"

// UserID returns the current user's ID.
func (c *Client) UserID(ctx context.Context, opts RequestOptions) (string, error) {
	user, err := c.GetUserMe(ctx, opts)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
