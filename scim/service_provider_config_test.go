package scim

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
)

func Test_clientGetServiceProviderConfig(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		isError    bool
		exp        *ServiceProviderConfig
	}{
		{
			statusCode: 200,
			isError:    false,
			exp: &ServiceProviderConfig{
				AuthenticationSchemes: []AuthenticationScheme{
					{
						Type:        "oauthbearertoken",
						Name:        "OAuth Bearer Token",
						Description: "Authentication Scheme using the OAuth Bearer Token Standard",
						SpecURL:     "http://tools.ietf.org/html/draft-ietf-oauth-v2-bearer-01",
						Primary:     true,
					},
				},
				Patch:          &PatchConfig{},
				Bulk:           &BulkConfig{},
				Filter:         &FilterConfig{},
				ChangePassword: &ChangePasswordConfig{},
				Sort:           &SortConfig{},
				Etag:           &EtagConfig{},
				XMLDataFormat:  &XMLDataFormatConfig{},
			},
			body: `{
	"authenticationSchemes": [{
		"type": "oauthbearertoken",
		"name": "OAuth Bearer Token",
		"description": "Authentication Scheme using the OAuth Bearer Token Standard",
		"specUrl": "http:\/\/tools.ietf.org\/html\/draft-ietf-oauth-v2-bearer-01",
		"primary": true
	}],
	"patch": {
		"supported": false
	},
	"bulk": {
		"supported": false,
		"maxOperations": 0,
		"maxPayloadSize": 0
	},
	"filter": {
		"supported": false,
		"maxResults": 0
	},
	"changePassword": {
		"supported": false
	},
	"sort": {
		"supported": false
	},
	"etag": {
		"supported": false
	},
	"xmlDataFormat": {
		"supported": false
	}
}`,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get("/scim/v1/ServiceProviderConfigs").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		cfg, resp, err := client.GetServiceProviderConfig(ctx)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, cfg)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, cfg)
		}
	}
}
