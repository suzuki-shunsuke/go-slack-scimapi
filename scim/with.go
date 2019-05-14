package scim

func (c *client) copy() *client {
	return &client{
		endpoint:       c.endpoint,
		token:          c.token,
		clientFn:       c.clientFn,
		isError:        c.isError,
		parseResp:      c.parseResp,
		parseErrorResp: c.parseErrorResp,
	}
}

// WithClientFn returns a shallow copy of c with its clientFn changed to fn.
func (c *client) WithClientFn(fn ClientFn) Client {
	if fn == nil {
		fn = clientFn
	}
	cl := c.copy()
	cl.clientFn = fn
	return cl
}

// WithParseResp returns a shallow copy of c with its parseResp changed to fn.
func (c *client) WithParseResp(fn ParseResp) Client {
	if fn == nil {
		fn = parseResp
	}
	cl := c.copy()
	cl.parseResp = fn
	return cl
}

// WithParseErrorResp returns a shallow copy of c with its parseErrorResp changed to fn.
func (c *client) WithParseErrorResp(fn ParseErrorResp) Client {
	if fn == nil {
		fn = parseErrorResp
	}
	cl := c.copy()
	cl.parseErrorResp = fn
	return cl
}

// WithIsError returns a shallow copy of c with its isError changed to fn.
func (c *client) WithIsError(fn IsError) Client {
	if fn == nil {
		fn = isError
	}
	cl := c.copy()
	cl.isError = fn
	return cl
}

// WithEndpoint returns a shallow copy of c with its endpoint changed to endpoint.
func (c *client) WithEndpoint(endpoint string) Client {
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	cl := c.copy()
	cl.endpoint = endpoint
	return cl
}
