package scim

// SetNewHTTPClient sets fn to c.
// If fn is nil, NewHTTPClientDefault is used.
func (c *Client) SetNewHTTPClient(fn NewHTTPClient) {
	if fn == nil {
		c.newHTTPClient = NewHTTPClientDefault
		return
	}
	c.newHTTPClient = fn
}

// SetParseResp sets fn to c.
// fn shouldn't close the response body.
// If fn is nil, ParseRespDefault is used.
func (c *Client) SetParseResp(fn ParseResp) {
	if fn == nil {
		c.parseResp = ParseRespDefault
		return
	}
	c.parseResp = fn
}

// SetParseErrorResp sets fn to c.
// fn shouldn't close the response body.
// If fn is nil, ParseErrorRespDefault is used.
func (c *Client) SetParseErrorResp(fn ParseErrorResp) {
	if fn == nil {
		c.parseErrorResp = ParseErrorRespDefault
		return
	}
	c.parseErrorResp = fn
}

// SetIsError sets fn to c.
// If fn is nil, IsErrorDefault is used.
func (c *Client) SetIsError(fn IsError) {
	if fn == nil {
		c.isError = IsErrorDefault
		return
	}
	c.isError = fn
}

// SetEndpoint sets fn to c.
// If endpoint is empty, DefaultEndpoint is used.
func (c *Client) SetEndpoint(endpoint string) {
	if endpoint == "" {
		c.endpoint = DefaultEndpoint
		return
	}
	c.endpoint = endpoint
}
