package weheat

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
)

const (
	DefaultTokenURL = "https://auth.weheat.nl/auth/realms/Weheat/protocol/openid-connect/token/"
)

var DefaultScopes = []string{"openid", "offline_access"}

// TokenSource provides access tokens for API requests.
type TokenSource interface {
	Token(ctx context.Context) (string, error)
}

// StaticToken uses a fixed access token.
type StaticToken string

func (s StaticToken) Token(_ context.Context) (string, error) {
	if s == "" {
		return "", errors.New("weheat: access token required")
	}
	return string(s), nil
}

// OAuthConfig defines OAuth refresh token settings.
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	RefreshToken string
	Scopes       []string
}

// OAuthTokenSource returns a TokenSource that refreshes access tokens using a refresh token.
func OAuthTokenSource(cfg OAuthConfig) (TokenSource, error) {
	if cfg.ClientID == "" {
		return nil, errors.New("weheat: client_id required")
	}
	if cfg.TokenURL == "" {
		cfg.TokenURL = DefaultTokenURL
	}
	if cfg.RefreshToken == "" {
		return nil, errors.New("weheat: refresh token required")
	}
	scopes := cfg.Scopes
	if len(scopes) == 0 {
		scopes = append([]string(nil), DefaultScopes...)
	}

	oauthCfg := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: cfg.TokenURL,
		},
		Scopes: scopes,
	}

	src := oauthCfg.TokenSource(context.Background(), &oauth2.Token{RefreshToken: cfg.RefreshToken})
	return oauthTokenSource{source: src}, nil
}

type oauthTokenSource struct {
	source oauth2.TokenSource
}

func (o oauthTokenSource) Token(ctx context.Context) (string, error) {
	token, err := o.source.Token()
	if err != nil {
		return "", err
	}
	if token.AccessToken == "" {
		return "", errors.New("weheat: empty access token")
	}
	return token.AccessToken, nil
}
