package scim

import (
	"net/http"
)

func (c *Client) copy() *Client {
	return &Client{
		endpoint:       c.endpoint,
		token:          c.token,
		httpClient:     c.httpClient,
		isError:        c.isError,
		parseResp:      c.parseResp,
		parseErrorResp: c.parseErrorResp,
	}
}

// WithHTTPClient returns a shallow copy of c with its httpClient changed to client.
// If client is nil, http.DefaultClient is used.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	cl := c.copy()
	if client == nil {
		client = http.DefaultClient
	}
	cl.httpClient = client
	return cl
}

// WithParseResp returns a shallow copy of c with its parseResp changed to fn.
// fn shouldn't close the response body.
// If fn is nil, ParseRespDefault is used.
func (c *Client) WithParseResp(fn ParseResp) *Client {
	if fn == nil {
		fn = ParseRespDefault
	}
	cl := c.copy()
	cl.parseResp = fn
	return cl
}

// WithParseErrorResp returns a shallow copy of c with its parseErrorResp changed to fn.
// fn shouldn't close the response body.
// If fn is nil, ParseErrorRespDefault is used.
func (c *Client) WithParseErrorResp(fn ParseErrorResp) *Client {
	if fn == nil {
		fn = ParseErrorRespDefault
	}
	cl := c.copy()
	cl.parseErrorResp = fn
	return cl
}

// WithIsError returns a shallow copy of c with its isError changed to fn.
// If fn is nil, IsErrorDefault is used.
func (c *Client) WithIsError(fn IsError) *Client {
	if fn == nil {
		fn = IsErrorDefault
	}
	cl := c.copy()
	cl.isError = fn
	return cl
}

// WithEndpoint returns a shallow copy of c with its endpoint changed to endpoint.
// If endpoint is empty, DefaultEndpoint is used.
func (c *Client) WithEndpoint(endpoint string) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	cl := c.copy()
	cl.endpoint = endpoint
	return cl
}
