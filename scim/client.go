package scim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
)

type (
	client struct {
		endpoint       string
		token          string
		clientFn       ClientFn
		isError        IsError
		parseResp      ParseResp
		parseErrorResp ParseErrorResp
	}

	// Client is a Slack SCIM API client.
	Client interface {
		GetUserSchemaResp(ctx context.Context) (*http.Response, error)
		GetUserSchema(ctx context.Context) (*Schema, *http.Response, error)

		GetGroupSchemaResp(ctx context.Context) (*http.Response, error)
		GetGroupSchema(ctx context.Context) (*Schema, *http.Response, error)

		GetServiceProviderConfig(context.Context) (*ServiceProviderConfig, *http.Response, error)
		GetServiceProviderConfigResp(context.Context) (*http.Response, error)

		GetGroupResp(ctx context.Context, id string) (*http.Response, error)
		GetGroup(ctx context.Context, id string) (*Group, *http.Response, error)

		CreateGroupResp(context.Context, *Group) (*http.Response, error)
		CreateGroup(context.Context, *Group) (*http.Response, error)

		PatchGroupResp(ctx context.Context, id string, group *Group) (*http.Response, error)
		PatchGroup(ctx context.Context, group *Group, id string) (*http.Response, error)

		PutGroupResp(ctx context.Context, id string, group *Group) (*http.Response, error)
		PutGroup(ctx context.Context, group *Group, id string) (*http.Response, error)

		DeleteGroupResp(ctx context.Context, id string) (*http.Response, error)
		DeleteGroup(ctx context.Context, id string) (*http.Response, error)

		GetGroupsResp(ctx context.Context, page *Pagination, filter string) (*http.Response, error)
		GetGroups(ctx context.Context, page *Pagination, filter string) (*Groups, *http.Response, error)

		GetUsersResp(ctx context.Context, page *Pagination, filter string) (*http.Response, error)
		GetUsers(ctx context.Context, page *Pagination, filter string) (*Users, *http.Response, error)

		GetUserResp(ctx context.Context, id string) (*http.Response, error)
		GetUser(ctx context.Context, id string) (*User, *http.Response, error)

		CreateUserResp(context.Context, *User) (*http.Response, error)
		CreateUser(context.Context, *User) (*http.Response, error)

		PatchUserResp(ctx context.Context, id string, user *User) (*http.Response, error)
		PatchUser(ctx context.Context, user *User, id string) (*http.Response, error)

		PutUserResp(ctx context.Context, id string, user *User) (*http.Response, error)
		PutUser(ctx context.Context, user *User, id string) (*http.Response, error)

		DeleteUserResp(ctx context.Context, id string) (*http.Response, error)
		DeleteUser(ctx context.Context, id string) (*http.Response, error)

		WithClientFn(ClientFn) Client
		WithParseResp(ParseResp) Client
		WithParseErrorResp(ParseErrorResp) Client
		WithIsError(IsError) Client
		WithEndpoint(endpoint string) Client
	}

	// ClientFn returns a new http.Client .
	ClientFn func() (*http.Client, error)
	// ParseResp parses a succeeded API response.
	ParseResp func(resp *http.Response, output interface{}) error
	// ParseErrorResp parses an API error response.
	ParseErrorResp func(resp *http.Response) error
	// IsError decides whether the request successes or not.
	IsError func(resp *http.Response) bool
)

var (
	defaultEndpoint = "https://api.slack.com/scim/v1"
)

// NewClient returns a new client.
func NewClient(token string) Client {
	return &client{
		token:          token,
		endpoint:       defaultEndpoint,
		clientFn:       clientFn,
		isError:        isError,
		parseResp:      parseResp,
		parseErrorResp: parseErrorResp,
	}
}

func (c *client) getResp(
	ctx context.Context, method, path string, body interface{}, query url.Values,
) (*http.Response, error) {
	endpoint, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}

	endpoint.Path = filepath.Join(endpoint.Path, path)
	endpoint.RawQuery = query.Encode()
	var req *http.Request
	if body == nil {
		req, err = http.NewRequest(method, endpoint.String(), nil)
	} else {
		reqBody := &bytes.Buffer{}
		if err := json.NewEncoder(reqBody).Encode(body); err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, endpoint.String(), reqBody)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)
	client, err := c.clientFn()
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (c *client) parseResponse(
	resp *http.Response, output interface{},
) error {
	if c.isError(resp) {
		return c.parseErrorResp(resp)
	}
	return c.parseResp(resp, output)
}
