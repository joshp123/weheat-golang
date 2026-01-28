package weheat

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	DefaultBaseURL   = "https://api.weheat.nl"
	DefaultUserAgent = "weheat-golang/0.1.0"
)

// ClientOption configures the Weheat client.
type ClientOption func(*Client) error

// WithBaseURL overrides the default API base URL.
func WithBaseURL(raw string) ClientOption {
	return func(c *Client) error {
		if raw == "" {
			return errors.New("weheat: base URL required")
		}
		parsed, err := url.Parse(raw)
		if err != nil {
			return err
		}
		c.baseURL = parsed
		return nil
	}
}

// WithHTTPClient sets the HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		if client == nil {
			return errors.New("weheat: http client required")
		}
		c.httpClient = client
		return nil
	}
}

// WithTokenSource sets the OAuth token source.
func WithTokenSource(source TokenSource) ClientOption {
	return func(c *Client) error {
		c.tokenSource = source
		return nil
	}
}

// WithUserAgent sets the User-Agent header.
func WithUserAgent(agent string) ClientOption {
	return func(c *Client) error {
		if agent == "" {
			return errors.New("weheat: user agent required")
		}
		c.userAgent = agent
		return nil
	}
}
