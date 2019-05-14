package scim

import (
	"context"
	"net/http"
)

func (c *client) GetUserSchemaResp(ctx context.Context) (*http.Response, error) {
	// GET /Schemas/Users
	return c.getResp(ctx, "GET", "/Schemas/Users", nil, nil)
}

func (c *client) GetUserSchema(ctx context.Context) (*Schema, *http.Response, error) {
	// GET /Schemas/Users
	resp, err := c.GetUserSchemaResp(ctx)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	schema := &Schema{}
	return schema, resp, c.parseResponse(resp, schema)
}
