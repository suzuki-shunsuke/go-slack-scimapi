package scim

import (
	"context"
	"net/http"
)

// GetGroupSchemaResp calls GET /Schemas/Groups API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) GetGroupSchemaResp(ctx context.Context) (*http.Response, error) {
	// GET /Schemas/Groups
	return c.getResp(ctx, "GET", "/Schemas/Groups", nil, nil)
}

// GetGroupSchema calls GET /Schemas/Groups API and returns a group schema.
// The returned response body is closed.
func (c *Client) GetGroupSchema(ctx context.Context) (*Schema, *http.Response, error) {
	// GET /Schemas/Groups
	resp, err := c.GetGroupSchemaResp(ctx)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	schema := &Schema{}
	return schema, resp, c.parseResponse(resp, schema)
}
