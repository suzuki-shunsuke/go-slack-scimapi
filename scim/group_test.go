package scim

import (
	"context"
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
)

const (
	dummyID = "XXXX"
)

var (
	testGroup = Group{
		Schemas: []string{
			"urn:scim:schemas:core:1.0",
		},
		ID: "XXXXXXXXX",
		Meta: &Meta{
			Created:  "2018-01-16T19:33:57-08:00",
			Location: "https://api.slack.com/scim/v1/Groups/XXXXXXXXX",
		},
		DisplayName: "Group Name",
		Members: []Member{
			{
				Value:   "YYYYY",
				Display: "First Last",
			},
		},
	}

	testGroupJSON = `{
  "schemas": [
    "urn:scim:schemas:core:1.0"
  ],
  "id": "XXXXXXXXX",
  "externalId": "",
  "meta": {
    "created": "2018-01-16T19:33:57-08:00",
    "location": "https://api.slack.com/scim/v1/Groups/XXXXXXXXX"
  },
  "displayName": "Group Name",
  "members": [
    {
      "value": "YYYYY",
      "display": "First Last"
    }
  ]
}`
)

func Test_clientGetGroups(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		isError    bool
		exp        *Groups
	}{
		{
			statusCode: 200,
			isError:    false,
			exp: &Groups{
				TotalResults: 1,
				Resources:    []Group{testGroup},
			},
			body: fmt.Sprintf(`{
  "totalResults": 1,
  "Resources": [
    %s
  ]
}`, testGroupJSON),
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get("/scim/v1/Groups").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		groups, resp, err := client.GetGroups(ctx, nil, "")
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, groups)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, groups)
		}
	}
}

func Test_clientGetGroup(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		id         string
		isError    bool
		exp        *Group
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			exp:        &testGroup,
			body:       testGroupJSON,
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			body:       testGroupJSON,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get(fmt.Sprintf("/scim/v1/Groups/%s", dummyID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		group, resp, err := client.GetGroup(ctx, d.id)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, group)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, group)
		}
	}
}

func Test_clientDeleteGroup(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		id         string
	}{
		{
			statusCode: 204,
			isError:    false,
			id:         dummyID,
		},
		{
			statusCode: 204,
			isError:    true,
			id:         "",
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Delete(fmt.Sprintf("/scim/v1/Groups/%s", dummyID)).
			MatchType("json").Reply(d.statusCode)
		resp, err := client.DeleteGroup(ctx, d.id)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}

func Test_clientCreateGroup(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		group      *Group
	}{
		{
			statusCode: 201,
			isError:    false,
			group:      &Group{},
		},
		{
			statusCode: 201,
			isError:    true,
			group:      nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Post("/scim/v1/Groups").
			MatchType("json").JSON(d.group).Reply(d.statusCode)
		resp, err := client.CreateGroup(ctx, d.group)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}

func Test_clientPutGroup(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		id         string
		group      *Group
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			group:      &Group{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			group:      &Group{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         dummyID,
			group:      nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Put(fmt.Sprintf("/scim/v1/Groups/%s", dummyID)).
			MatchType("json").JSON(d.group).Reply(d.statusCode)
		resp, err := client.PutGroup(ctx, d.id, d.group)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}

func Test_clientPatchGroup(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		id         string
		group      *Group
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			group:      &Group{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			group:      &Group{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         dummyID,
			group:      nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Patch(fmt.Sprintf("/scim/v1/Groups/%s", dummyID)).
			MatchType("json").JSON(d.group).Reply(d.statusCode)
		resp, err := client.PatchGroup(ctx, d.id, d.group)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}
