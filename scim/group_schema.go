package scim

import (
	"context"
	"net/http"
)

func (c *Client) GetGroupSchemaResp(ctx context.Context) (*http.Response, error) {
	// GET /Schemas/Groups
	return c.getResp(ctx, "GET", "/Schemas/Groups", nil, nil)
}

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
