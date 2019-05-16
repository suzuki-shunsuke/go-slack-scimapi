package scim

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// User is a user.
	// https://api.slack.com/scim#users
	User struct {
		Active                        bool                           `json:"active,omitempty"`
		ID                            string                         `json:"id,omitempty"`
		ExternalID                    string                         `json:"externalId,omitempty"`
		UserName                      string                         `json:"userName,omitempty"`
		NickName                      string                         `json:"nickName,omitempty"`
		ProfileURL                    string                         `json:"profileUrl,omitempty"`
		DisplayName                   string                         `json:"displayName,omitempty"`
		UserType                      string                         `json:"userType,omitempty"`
		Title                         string                         `json:"title,omitempty"`
		PreferredLanguage             string                         `json:"preferredLanguage,omitempty"`
		Locale                        string                         `json:"locale,omitempty"`
		Timezone                      string                         `json:"timezone,omitempty"`
		Password                      string                         `json:"password,omitempty"`
		Name                          *Name                          `json:"name,omitempty"`
		Meta                          *Meta                          `json:"meta,omitempty"`
		Emails                        []Email                        `json:"emails,omitempty"`
		Addresses                     []Address                      `json:"addresses,omitempty"`
		PhoneNumbers                  []PhoneNumber                  `json:"phoneNumbers,omitempty"`
		Roles                         []Role                         `json:"roles,omitempty"`
		Photos                        []Photo                        `json:"photos,omitempty"`
		Groups                        []Group                        `json:"groups,omitempty"`
		Schemas                       []string                       `json:"schemas"`
		EnterpriseUserSchemaExtension *EnterpriseUserSchemaExtension `json:"urn:scim:schemas:extension:enterprise:1.0,omitempty"`
	}

	// EnterpriseUserSchemaExtension is SCIM Enterprise User Schema Extension.
	EnterpriseUserSchemaExtension struct {
		EmployeeNumber string   `json:"employeeNumber,omitempty"`
		CostCenter     string   `json:"costCenter,omitempty"`
		Organization   string   `json:"organization,omitempty"`
		Division       string   `json:"division,omitempty"`
		Department     string   `json:"department,omitempty"`
		Manager        *Manager `json:"manager,omitempty"`
	}

	// Manager is a user's manager.
	Manager struct {
		ManagerID   string `json:"managerId,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
	}

	// Email is an email address for a user.
	Email struct {
		Value   string `json:"value,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	}

	// Address is a physical mailing address for a user.
	Address struct {
		StreetAddress string `json:"streetAddress,omitempty"`
		Locality      string `json:"locality,omitempty"`
		Region        string `json:"region,omitempty"`
		PostalCode    string `json:"postalCode,omitempty"`
		Country       string `json:"country,omitempty"`
		Primary       bool   `json:"primary,omitempty"`
	}

	// PhoneNumber is a phone number for a user.
	PhoneNumber struct {
		Value   string `json:"value,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	}

	// Photo is a URL of a photo of a User.
	Photo struct {
		Value   string `json:"value,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	}

	// Role is a role for a user that collectively represent who the user is; e.g., 'Student', "Faculty".
	Role struct {
		Value   string `json:"value,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	}

	// Name is a component of a user's real name.
	Name struct {
		FamilyName      string `json:"familyName,omitempty"`
		GivenName       string `json:"givenName,omitempty"`
		HonorificPrefix string `json:"honorificPrefix,omitempty"`
	}

	// Users is a response body of GET users API .
	Users struct {
		TotalResults int      `json:"totalResults"`
		ItemPerPage  int      `json:"itemPerPage"`
		StartIndex   int      `json:"startIndex"`
		Schemas      []string `json:"schemas"`
		Resources    []User
	}

	userGroup struct {
		Value   string `json:"value"`
		Display string `json:"display"`
	}
)

// GetUsersResp calls GET /Users API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) GetUsersResp(ctx context.Context, page *Pagination, filter string) (*http.Response, error) {
	// GET /Users
	query := url.Values{}
	if filter != "" {
		query.Add("filter", filter)
	}
	setPageToQuery(page, query)
	return c.getResp(ctx, "GET", "/Users", nil, query)
}

// GetUsers calls GET /Users API and returns users.
// The returned response body is closed.
func (c *Client) GetUsers(
	ctx context.Context, page *Pagination, filter string,
) (*Users, *http.Response, error) {
	// GET /Users
	resp, err := c.GetUsersResp(ctx, page, filter)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	users := &Users{}
	return users, resp, c.parseResponse(resp, users)
}

// GetUserResp calls GET /Users/{id} API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) GetUserResp(ctx context.Context, id string) (*http.Response, error) {
	// GET /Users/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.getResp(ctx, "GET", fmt.Sprintf("/Users/%s", id), nil, nil)
}

// GetUser calls GET /Users/{id} API and returns a user.
// The returned response body is closed.
func (c *Client) GetUser(
	ctx context.Context, id string,
) (*User, *http.Response, error) {
	// GET /Users/{id}
	resp, err := c.GetUserResp(ctx, id)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	user := &User{}
	return user, resp, c.parseResponse(resp, user)
}

// CreateUserResp calls POST /Users API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) CreateUserResp(ctx context.Context, user *User) (*http.Response, error) {
	// POST /Users
	if user == nil {
		return nil, fmt.Errorf("user is required")
	}
	return c.getResp(ctx, "POST", "/Users", user, nil)
}

// CreateGroup calls POST /Users API and returns a created user.
// The returned response body is closed.
func (c *Client) CreateUser(ctx context.Context, user *User) (*User, *http.Response, error) {
	// POST /Users
	resp, err := c.CreateUserResp(ctx, user)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	u := &User{}
	return u, resp, c.parseResponse(resp, u)
}

// PatchUserResp calls PATCH /Users/{id} API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) PatchUserResp(ctx context.Context, id string, user *UserPatch) (*http.Response, error) {
	// PATCH /Users/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if user == nil {
		return nil, fmt.Errorf("user is required")
	}
	return c.getResp(ctx, "PATCH", fmt.Sprintf("/Users/%s", id), user, nil)
}

// PatchUser calls PATCH /Users/{id} API and returns a updated user.
// The returned response body is closed.
func (c *Client) PatchUser(ctx context.Context, id string, user *UserPatch) (*User, *http.Response, error) {
	// PATCH /Users/{id}
	resp, err := c.PatchUserResp(ctx, id, user)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	u := &User{}
	return u, resp, c.parseResponse(resp, u)
}

// PutUserResp calls PUT /Users/{id} API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) PutUserResp(ctx context.Context, id string, user *User) (*http.Response, error) {
	// PUT /Users/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if user == nil {
		return nil, fmt.Errorf("user is required")
	}
	return c.getResp(ctx, "PUT", fmt.Sprintf("/Users/%s", id), user, nil)
}

// PutUser calls PUT /Users/{id} API and returns a updated user.
// The returned response body is closed.
func (c *Client) PutUser(ctx context.Context, id string, user *User) (*User, *http.Response, error) {
	// PUT /Users/{id}
	resp, err := c.PutUserResp(ctx, id, user)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	u := &User{}
	return u, resp, c.parseResponse(resp, u)
}

// DeleteUserResp calls DELETE /Users/{id} API and returns a HTTP response.
// If the returned error is nil, the returned response isn't nil and you have to close the response body.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *Client) DeleteUserResp(ctx context.Context, id string) (*http.Response, error) {
	// DELETE /Users/{id}
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return c.getResp(ctx, "DELETE", fmt.Sprintf("/Users/%s", id), nil, nil)
}

// DeleteUser calls DELETE /Users/{id} API.
// The returned response body is closed.
func (c *Client) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	// DELETE /Users/{id}
	resp, err := c.DeleteUserResp(ctx, id)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, c.parseResponse(resp, nil)
}

// UnmarshalJSON implements json.Unmarshaler .
func (user *User) UnmarshalJSON(b []byte) error {
	type alias User
	a := struct {
		Groups json.RawMessage `json:"groups"`
		*alias
	}{
		alias: (*alias)(user),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if len(a.Groups) == 0 {
		return nil
	}

	list := []userGroup{}
	if err := json.Unmarshal(a.Groups, &list); err != nil {
		return err
	}
	groups := []Group{}
	for _, g := range list {
		groups = append(groups, Group{
			ID:          g.Value,
			DisplayName: g.Display,
		})
	}
	user.Groups = groups
	return nil
}

// MarshalJSON implements json.Marshaler .
func (user *User) MarshalJSON() ([]byte, error) {
	list := make([]userGroup, len(user.Groups))
	for i, group := range user.Groups {
		list[i] = userGroup{
			Value:   group.ID,
			Display: group.DisplayName,
		}
	}
	type alias User
	a := struct {
		Groups []userGroup `json:"groups"`
		*alias
	}{
		Groups: list,
		alias:  (*alias)(user),
	}
	return json.Marshal(a)
}
