package scim

import (
	"context"
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
)

var (
	testUser = User{
		Schemas: []string{
			"urn:scim:schemas:core:1.0",
			"urn:scim:schemas:extension:enterprise:1.0",
		},
		ID:         "XXXXXXXXX",
		ExternalID: "",
		Meta: &Meta{
			Created:  "2018-01-16T19:33:57-08:00",
			Location: "https://api.slack.com/scim/v1/Users/XXXXXXXXX",
		},
		UserName: "other_username",
		NickName: "slack_username",
		Name: &Name{
			GivenName:       "First",
			FamilyName:      "Last",
			HonorificPrefix: "Ms.",
		},
		DisplayName: "First Last",
		ProfileURL:  "https://login.example.com/slack_username",
		Title:       "Tour Guide",
		Timezone:    "America/Chicago",
		Active:      true,
		Emails: []Email{
			{
				Value:   "some@example.com",
				Type:    "work",
				Primary: true,
			},
			{
				Value: "some_other@example.com",
				Type:  "home",
			},
		},
		Photos: []Photo{
			{
				Value: "https://photos.example.com/profilephoto.jpg",
				Type:  "photo",
			},
		},
		UserType: "Employee",
		Addresses: []Address{
			{
				StreetAddress: "1060 W Addison St",
				Locality:      "Chicago",
				Region:        "IL",
				PostalCode:    "60613",
				Country:       "USA",
				Primary:       true,
			},
		},
		PhoneNumbers: []PhoneNumber{
			{
				Value: "555-555-5555",
				Type:  "work",
			},
			{
				Value: "555-555-4444",
				Type:  "mobile",
			},
		},
		Roles: []Role{
			{
				Value:   "Tech Lead",
				Primary: true,
			},
		},
		EnterpriseUserSchemaExtension: &EnterpriseUserSchemaExtension{
			EmployeeNumber: "701984",
			CostCenter:     "4130",
			Organization:   "Chicago Cubs",
			Division:       "Wrigley Field",
			Department:     "Tour Operations",
			Manager: &Manager{
				ManagerID: "U0XE15NHQ",
			},
		},
		PreferredLanguage: "en_US",
		Locale:            "en_US",
		Groups: []Group{
			{
				ID:          "YYYYYYYYY",
				DisplayName: "Group name",
			},
		},
	}

	testUserJSON = `{
  "schemas": [
    "urn:scim:schemas:core:1.0",
    "urn:scim:schemas:extension:enterprise:1.0"
  ],
  "id": "XXXXXXXXX",
  "externalId": "",
  "userType": "Employee",
  "meta": {
    "created": "2018-01-16T19:33:57-08:00",
    "location": "https://api.slack.com/scim/v1/Users/XXXXXXXXX"
  },
  "userName": "other_username",
  "nickName": "slack_username",
  "name": {
    "givenName": "First",
    "familyName": "Last",
    "honorificPrefix": "Ms."
  },
  "displayName": "First Last",
  "profileUrl": "https://login.example.com/slack_username",
  "title": "Tour Guide",
  "timezone": "America/Chicago",
  "active": true,
  "emails": [
    {
      "value": "some@example.com",
      "type": "work",
      "primary": true
    },
    {
      "value": "some_other@example.com",
      "type": "home"
    }
  ],
  "photos": [
    {
      "value": "https://photos.example.com/profilephoto.jpg",
      "type": "photo"
    }
  ],
  "addresses": [
    {
      "streetAddress": "1060 W Addison St",
      "locality": "Chicago",
      "region": "IL",
      "postalCode": "60613",
      "country": "USA",
      "primary": true
    }
  ],
  "phoneNumbers": [
    {
      "value": "555-555-5555",
      "type": "work"
    },
    {
      "value": "555-555-4444",
      "type": "mobile"
    }
  ],
  "userType": "Employee",
  "roles": [
    {
      "value": "Tech Lead",
      "primary": true
    }
  ],
  "preferredLanguage": "en_US",
  "locale": "en_US",
  "urn:scim:schemas:extension:enterprise:1.0": {
    "employeeNumber": "701984",
    "costCenter": "4130",
    "organization": "Chicago Cubs",
    "division": "Wrigley Field",
    "department": "Tour Operations",
    "manager": {
      "managerId": "U0XE15NHQ"
    }
  },
  "groups": [
    {
      "value": "YYYYYYYYY",
      "display": "Group name"
    }
  ]
}`
)

func TestClientGetUsers(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		isError    bool
		exp        *Users
	}{
		{
			statusCode: 200,
			isError:    false,
			exp: &Users{
				TotalResults: 1,
				Resources:    []User{testUser},
			},
			body: fmt.Sprintf(`{
  "totalResults": 1,
  "Resources": [
    %s
  ]
}`, testUserJSON),
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get("/scim/v1/Users").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		users, resp, err := client.GetUsers(ctx, nil, `email eq "some@example.com"`)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, users)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, users)
		}
	}
}

func TestClientGetUser(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		id         string
		isError    bool
		exp        *User
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			exp:        &testUser,
			body:       testUserJSON,
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			exp:        nil,
			body:       testUserJSON,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get(fmt.Sprintf("/scim/v1/Users/%s", dummyID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		user, resp, err := client.GetUser(ctx, d.id)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, user)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, user)
		}
	}
}

func TestClientCreateUser(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		user       *User
		resp       *User
	}{
		{
			statusCode: 201,
			isError:    false,
			user:       &User{},
			resp:       &testUser,
		},
		{
			statusCode: 201,
			isError:    true,
			user:       nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Post("/scim/v1/Users").
			MatchType("json").JSON(d.user).Reply(d.statusCode).JSON(d.resp)
		user, resp, err := client.CreateUser(ctx, d.user)
		gock.Off()
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.resp != nil {
			require.Equal(t, d.resp, user)
		}
	}
}

func TestClientPutUser(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		id         string
		user       *User
		resp       *User
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			user:       &testUser,
			resp:       &testUser,
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			user:       &User{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         dummyID,
			user:       nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Put(fmt.Sprintf("/scim/v1/Users/%s", dummyID)).
			MatchType("json").JSON(d.user).Reply(d.statusCode).JSON(d.resp)
		user, resp, err := client.PutUser(ctx, d.id, d.user)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.resp != nil {
			require.Equal(t, d.resp, user)
		}
	}
}

func TestClientPatchUser(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		isError    bool
		id         string
		user       *UserPatch
	}{
		{
			statusCode: 200,
			isError:    false,
			id:         dummyID,
			user:       &UserPatch{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         "",
			user:       &UserPatch{},
		},
		{
			statusCode: 200,
			isError:    true,
			id:         dummyID,
			user:       nil,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Patch(fmt.Sprintf("/scim/v1/Users/%s", dummyID)).
			MatchType("json").JSON(d.user).Reply(d.statusCode)
		resp, err := client.PatchUser(ctx, d.id, d.user)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}

func TestClientDeleteUser(t *testing.T) {
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
			Delete(fmt.Sprintf("/scim/v1/Users/%s", dummyID)).
			MatchType("json").Reply(d.statusCode)
		resp, err := client.DeleteUser(ctx, d.id)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, d.statusCode, resp.StatusCode)
	}
}
