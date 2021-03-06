package scim

import (
	"context"
	"net/http"
)

// GetUserSchemaResp calls GET /Schemas/Users API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) GetUserSchemaResp(ctx context.Context) (*http.Response, error) {
	// GET /Schemas/Users
	return c.getResp(ctx, "GET", "/Schemas/Users", nil, nil)
}

// GetUserSchema calls GET /Schemas/Users API and returns a user schema.
// The returned response body is closed.
func (c *Client) GetUserSchema(ctx context.Context) (*Schema, *http.Response, error) {
	// GET /Schemas/Users
	resp, err := c.GetUserSchemaResp(ctx)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	schema := &Schema{}
	return schema, resp, c.parseResponse(resp, schema)
}
