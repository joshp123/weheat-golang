package weheat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.000000-0700"

// Client talks to the Weheat API.
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	tokenSource TokenSource
	userAgent   string
}

// NewClient creates a new client with optional overrides.
func NewClient(opts ...ClientOption) (*Client, error) {
	base, err := url.Parse(DefaultBaseURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		baseURL:    base,
		httpClient: &http.Client{Timeout: 15 * time.Second},
		userAgent:  DefaultUserAgent,
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// GetUserMe fetches the current user profile.
func (c *Client) GetUserMe(ctx context.Context, opts RequestOptions) (*ReadUserMe, error) {
	var out ReadUserMe
	if err := c.doJSON(ctx, http.MethodGet, "/api/v1/users/me", nil, opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListHeatPumps returns a paged list of heat pumps.
func (c *Client) ListHeatPumps(ctx context.Context, params ListHeatPumpsParams) (*ReadAllHeatPumpPagedResponse, error) {
	query := url.Values{}
	if params.Page != nil {
		query.Set("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.PageSize != nil {
		query.Set("pageSize", fmt.Sprintf("%d", *params.PageSize))
	}
	if len(params.Models) > 0 {
		for _, model := range params.Models {
			query.Add("Model", fmt.Sprintf("%d", model))
		}
	}
	if params.OrganisationID != "" {
		query.Set("OrganisationId", params.OrganisationID)
	}
	if params.Search != "" {
		query.Set("Search", params.Search)
	}
	if params.State != nil {
		query.Set("State", fmt.Sprintf("%d", *params.State))
	}

	var out ReadAllHeatPumpPagedResponse
	if err := c.doJSON(ctx, http.MethodGet, "/api/v1/heat-pumps", query, params.RequestOptions, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHeatPump returns a single heat pump by ID.
func (c *Client) GetHeatPump(ctx context.Context, heatPumpID string, opts RequestOptions) (*ReadHeatPump, error) {
	path := fmt.Sprintf("/api/v1/heat-pumps/%s", url.PathEscape(heatPumpID))
	var out ReadHeatPump
	if err := c.doJSON(ctx, http.MethodGet, path, nil, opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetLatestLog returns the most recent log entry for a heat pump.
func (c *Client) GetLatestLog(ctx context.Context, heatPumpID string, opts RequestOptions) (*RawHeatPumpLog, error) {
	path := fmt.Sprintf("/api/v1/heat-pumps/%s/logs/latest", url.PathEscape(heatPumpID))
	var out RawHeatPumpLog
	if err := c.doJSON(ctx, http.MethodGet, path, nil, opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRawLogs returns raw log entries for a heat pump.
func (c *Client) GetRawLogs(ctx context.Context, heatPumpID string, query LogQuery) ([]RawHeatPumpLog, error) {
	path := fmt.Sprintf("/api/v1/heat-pumps/%s/logs/raw", url.PathEscape(heatPumpID))
	values := url.Values{}
	applyLogQuery(values, query)

	var out []RawHeatPumpLog
	if err := c.doJSON(ctx, http.MethodGet, path, values, query.RequestOptions, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetLogs returns aggregated log views for a heat pump.
func (c *Client) GetLogs(ctx context.Context, heatPumpID string, query LogQuery) ([]HeatPumpLogView, error) {
	path := fmt.Sprintf("/api/v1/heat-pumps/%s/logs", url.PathEscape(heatPumpID))
	values := url.Values{}
	applyLogQuery(values, query)

	var out []HeatPumpLogView
	if err := c.doJSON(ctx, http.MethodGet, path, values, query.RequestOptions, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetEnergyLogs returns aggregated energy logs for a heat pump.
func (c *Client) GetEnergyLogs(ctx context.Context, heatPumpID string, query EnergyLogQuery) ([]EnergyView, error) {
	path := fmt.Sprintf("/api/v1/energy-logs/%s", url.PathEscape(heatPumpID))
	values := url.Values{}
	applyEnergyQuery(values, query)

	var out []EnergyView
	if err := c.doJSON(ctx, http.MethodGet, path, values, query.RequestOptions, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetEnergyTotals returns total energy aggregates for a heat pump.
func (c *Client) GetEnergyTotals(ctx context.Context, heatPumpID string, opts RequestOptions) (*TotalEnergyAggregate, error) {
	path := fmt.Sprintf("/api/v1/energy-logs/%s/total", url.PathEscape(heatPumpID))
	var out TotalEnergyAggregate
	if err := c.doJSON(ctx, http.MethodGet, path, nil, opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func applyLogQuery(values url.Values, query LogQuery) {
	if query.StartTime != nil {
		values.Set("startTime", query.StartTime.UTC().Format(timeFormat))
	}
	if query.EndTime != nil {
		values.Set("endTime", query.EndTime.UTC().Format(timeFormat))
	}
	if query.Interval != "" {
		values.Set("interval", string(query.Interval))
	}
}

func applyEnergyQuery(values url.Values, query EnergyLogQuery) {
	if query.StartTime != nil {
		values.Set("startTime", query.StartTime.UTC().Format(timeFormat))
	}
	if query.EndTime != nil {
		values.Set("endTime", query.EndTime.UTC().Format(timeFormat))
	}
	if query.Interval != "" {
		values.Set("interval", string(query.Interval))
	}
}

func (c *Client) doJSON(ctx context.Context, method string, path string, query url.Values, opts RequestOptions, out any) error {
	headers := map[string]string{
		"Accept": "application/json, text/json, text/plain",
	}
	applyRequestOptions(headers, opts)

	req, err := c.newRequest(ctx, method, path, query, headers)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{StatusCode: resp.StatusCode, Body: body}
	}

	if out == nil || len(body) == 0 {
		return nil
	}

	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if err := dec.Decode(out); err != nil {
		return err
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, method string, path string, query url.Values, headers map[string]string) (*http.Request, error) {
	endpoint, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		endpoint.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	for key, value := range headers {
		if value == "" {
			continue
		}
		req.Header.Set(key, value)
	}

	if c.tokenSource != nil {
		token, err := c.tokenSource.Token(ctx)
		if err != nil {
			return nil, err
		}
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+strings.TrimSpace(token))
		}
	}

	return req, nil
}

func applyRequestOptions(headers map[string]string, opts RequestOptions) {
	if opts.XVersion != "" {
		headers["x-version"] = opts.XVersion
	}
	if opts.XBackendVersion != "" {
		headers["x-backend-version"] = opts.XBackendVersion
	}
}
