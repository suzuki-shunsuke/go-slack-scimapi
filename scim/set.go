package scim

// SetNewHTTPClient sets fn to c.
func (c *client) SetNewHTTPClient(fn NewHTTPClient) {
	if fn == nil {
		c.newHTTPClient = NewHTTPClientDefault
		return
	}
	c.newHTTPClient = fn
}

// SetParseResp sets fn to c.
func (c *client) SetParseResp(fn ParseResp) {
	if fn == nil {
		c.parseResp = ParseRespDefault
		return
	}
	c.parseResp = fn
}

// SetParseErrorResp sets fn to c.
func (c *client) SetParseErrorResp(fn ParseErrorResp) {
	if fn == nil {
		c.parseErrorResp = ParseErrorRespDefault
		return
	}
	c.parseErrorResp = fn
}

// SetIsError sets fn to c.
func (c *client) SetIsError(fn IsError) {
	if fn == nil {
		c.isError = IsErrorDefault
		return
	}
	c.isError = fn
}

// SetEndpoint sets fn to c.
func (c *client) SetEndpoint(endpoint string) {
	if endpoint == "" {
		c.endpoint = DefaultEndpoint
		return
	}
	c.endpoint = endpoint
}
