package scim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
)

type (
	// Client is a Slack SCIM API client.
	// Client should be created by the function NewClient .
	Client struct {
		endpoint       string
		token          string
		newHTTPClient  NewHTTPClient
		isError        IsError
		parseResp      ParseResp
		parseErrorResp ParseErrorResp
	}

	// NewHTTPClient returns a new http.Client .
	NewHTTPClient func() (*http.Client, error)
	// ParseResp parses a succeeded API response.
	ParseResp func(resp *http.Response, output interface{}) error
	// ParseErrorResp parses an API error response.
	ParseErrorResp func(resp *http.Response) error
	// IsError decides whether the request successes or not.
	IsError func(resp *http.Response) bool
)

var (
	// DefaultEndpoint is the default Slack SCIM API endpoint.
	DefaultEndpoint = "https://api.slack.com/scim/v1"
)

// NewClient returns a new client.
func NewClient(token string) *Client {
	return &Client{
		token:          token,
		endpoint:       DefaultEndpoint,
		newHTTPClient:  NewHTTPClientDefault,
		isError:        IsErrorDefault,
		parseResp:      ParseRespDefault,
		parseErrorResp: ParseErrorRespDefault,
	}
}

func (c *Client) getResp(
	ctx context.Context, method, path string, body interface{}, query url.Values,
) (*http.Response, error) {
	endpoint, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}

	endpoint.Path = filepath.Join(endpoint.Path, path)
	endpoint.RawQuery = query.Encode()
	var req *http.Request
	if body == nil {
		req, err = http.NewRequest(method, endpoint.String(), nil)
	} else {
		reqBody := &bytes.Buffer{}
		if err := json.NewEncoder(reqBody).Encode(body); err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, endpoint.String(), reqBody)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)
	client, err := c.newHTTPClient()
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (c *Client) parseResponse(
	resp *http.Response, output interface{},
) error {
	if c.isError(resp) {
		return c.parseErrorResp(resp)
	}
	return c.parseResp(resp, output)
}
