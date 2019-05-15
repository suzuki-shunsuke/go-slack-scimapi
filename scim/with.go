package scim

func (c *client) copy() *client {
	return &client{
		endpoint:       c.endpoint,
		token:          c.token,
		newHTTPClient:  c.newHTTPClient,
		isError:        c.isError,
		parseResp:      c.parseResp,
		parseErrorResp: c.parseErrorResp,
	}
}

// WithNewHTTPClient returns a shallow copy of c with its nweHTTPClient changed to fn.
func (c *client) WithNewHTTPClient(fn NewHTTPClient) Client {
	if fn == nil {
		fn = NewHTTPClientDefault
	}
	cl := c.copy()
	cl.newHTTPClient = fn
	return cl
}

// WithParseResp returns a shallow copy of c with its parseResp changed to fn.
func (c *client) WithParseResp(fn ParseResp) Client {
	if fn == nil {
		fn = ParseRespDefault
	}
	cl := c.copy()
	cl.parseResp = fn
	return cl
}

// WithParseErrorResp returns a shallow copy of c with its parseErrorResp changed to fn.
func (c *client) WithParseErrorResp(fn ParseErrorResp) Client {
	if fn == nil {
		fn = ParseErrorRespDefault
	}
	cl := c.copy()
	cl.parseErrorResp = fn
	return cl
}

// WithIsError returns a shallow copy of c with its isError changed to fn.
func (c *client) WithIsError(fn IsError) Client {
	if fn == nil {
		fn = IsErrorDefault
	}
	cl := c.copy()
	cl.isError = fn
	return cl
}

// WithEndpoint returns a shallow copy of c with its endpoint changed to endpoint.
func (c *client) WithEndpoint(endpoint string) Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	cl := c.copy()
	cl.endpoint = endpoint
	return cl
}
