package scim

func (c *Client) copy() *Client {
	return &Client{
		endpoint:       c.endpoint,
		token:          c.token,
		newHTTPClient:  c.newHTTPClient,
		isError:        c.isError,
		parseResp:      c.parseResp,
		parseErrorResp: c.parseErrorResp,
	}
}

// WithNewHTTPClient returns a shallow copy of c with its nweHTTPClient changed to fn.
func (c *Client) WithNewHTTPClient(fn NewHTTPClient) *Client {
	if fn == nil {
		fn = NewHTTPClientDefault
	}
	cl := c.copy()
	cl.newHTTPClient = fn
	return cl
}

// WithParseResp returns a shallow copy of c with its parseResp changed to fn.
func (c *Client) WithParseResp(fn ParseResp) *Client {
	if fn == nil {
		fn = ParseRespDefault
	}
	cl := c.copy()
	cl.parseResp = fn
	return cl
}

// WithParseErrorResp returns a shallow copy of c with its parseErrorResp changed to fn.
func (c *Client) WithParseErrorResp(fn ParseErrorResp) *Client {
	if fn == nil {
		fn = ParseErrorRespDefault
	}
	cl := c.copy()
	cl.parseErrorResp = fn
	return cl
}

// WithIsError returns a shallow copy of c with its isError changed to fn.
func (c *Client) WithIsError(fn IsError) *Client {
	if fn == nil {
		fn = IsErrorDefault
	}
	cl := c.copy()
	cl.isError = fn
	return cl
}

// WithEndpoint returns a shallow copy of c with its endpoint changed to endpoint.
func (c *Client) WithEndpoint(endpoint string) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	cl := c.copy()
	cl.endpoint = endpoint
	return cl
}
