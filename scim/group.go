package scim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type (
	Group struct {
		ID          string   `json:"id"`
		DisplayName string   `json:"displayName"`
		Members     []Member `json:"members"`
		Schemas     []string `json:"schemas"`
		Meta        *Meta    `json:"meta"`
	}

	Groups struct {
		TotalResults int      `json:"totalResults"`
		ItemPerPage  int      `json:"itemPerPage"`
		StartIndex   int      `json:"startIndex"`
		Schemas      []string `json:"schemas"`
		Resources    []Group
	}

	// Member is member of the group.
	Member struct {
		Value   string `json:"value"`
		Display string `json:"display"`
	}
)

// GetGroupsResp sends GET groups API and returns an HTTP response.
func (c *client) GetGroupsResp(ctx context.Context, page *Pagination, filter string) (*http.Response, error) {
	// GET /Groups
	query := url.Values{}
	if filter != "" {
		query.Add("filter", filter)
	}
	setPageToQuery(page, query)
	return c.getResp(ctx, "GET", "/Groups", nil, query)
}

// GetGroups sends GET groups API and returns groups.
func (c *client) GetGroups(
	ctx context.Context, page *Pagination, filter string,
) (*Groups, *http.Response, error) {
	// GET /Groups
	resp, err := c.GetGroupsResp(ctx, page, filter)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	groups := &Groups{}
	return groups, resp, c.parseResponse(resp, groups)
}

// GetGroupResp sends GET a group API and returns an HTTP response.
func (c *client) GetGroupResp(ctx context.Context, id string) (*http.Response, error) {
	// GET /Groups/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.getResp(ctx, "GET", fmt.Sprintf("/Groups/%s", id), nil, nil)
}

// GetGroup sends GET a group API and returns a group.
func (c *client) GetGroup(
	ctx context.Context, id string,
) (*Group, *http.Response, error) {
	// GET /Groups/{id}
	resp, err := c.GetGroupResp(ctx, id)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	group := &Group{}
	return group, resp, c.parseResponse(resp, group)
}

// CreateGroupResp sends POST a group API and returns an HTTP response.
func (c *client) CreateGroupResp(ctx context.Context, group *Group) (*http.Response, error) {
	// POST /Groups
	if group == nil {
		return nil, fmt.Errorf("group is required")
	}
	return c.getResp(ctx, "POST", "/Groups", group, nil)
}

// CreateGroup sends POST a group API.
func (c *client) CreateGroup(ctx context.Context, group *Group) (*http.Response, error) {
	// POST /Groups
	resp, err := c.CreateGroupResp(ctx, group)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, c.parseResponse(resp, nil)
}

// PatchGroupResp sends PATCH a group API and returns an HTTP response.
func (c *client) PatchGroupResp(ctx context.Context, id string, group *Group) (*http.Response, error) {
	// PATCH /Groups/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if group == nil {
		return nil, fmt.Errorf("group is required")
	}
	return c.getResp(ctx, "PATCH", fmt.Sprintf("/Groups/%s", id), group, nil)
}

// PatchGroup sends PATCH a group API.
func (c *client) PatchGroup(ctx context.Context, id string, group *Group) (*http.Response, error) {
	// PATCH /Groups/{id}
	resp, err := c.PatchGroupResp(ctx, id, group)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, c.parseResponse(resp, nil)
}

// PutGroupResp sends PUT a group API and returns an HTTP response.
func (c *client) PutGroupResp(ctx context.Context, id string, group *Group) (*http.Response, error) {
	// PUT /Groups/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if group == nil {
		return nil, fmt.Errorf("group is required")
	}
	return c.getResp(ctx, "PUT", fmt.Sprintf("/Groups/%s", id), group, nil)
}

// PutGroup sends PUT a group API.
func (c *client) PutGroup(ctx context.Context, group *Group, id string) (*http.Response, error) {
	// PUT /Groups/{id}
	resp, err := c.PutGroupResp(ctx, id, group)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, c.parseResponse(resp, nil)
}

// DeleteGroupResp sends DELETE a group API and returns an HTTP response.
func (c *client) DeleteGroupResp(ctx context.Context, id string) (*http.Response, error) {
	// DELETE /Groups/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.getResp(ctx, "DELETE", fmt.Sprintf("/Groups/%s", id), nil, nil)
}

// DeleteGroup sends DELETE a group API.
func (c *client) DeleteGroup(ctx context.Context, id string) (*http.Response, error) {
	// DELETE /Groups/{id}
	resp, err := c.DeleteGroupResp(ctx, id)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, c.parseResponse(resp, nil)
}
