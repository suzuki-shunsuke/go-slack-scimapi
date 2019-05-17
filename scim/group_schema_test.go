package scim

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
)

func TestClientGetGroupSchema(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		isError    bool
		exp        *Schema
	}{
		{
			statusCode: 200,
			isError:    false,
			exp: &Schema{
				Name:        "Group",
				Description: "Core Group",
				Schema: []string{
					"urn:scim:schemas:core:1.0",
				},
				Endpoint: "/Groups",
				Attributes: []Attribute{
					{
						Name:        "id",
						Type:        "string",
						Description: "Unique identifier for the SCIM resource as defined by the Service Provider. Each representation of the resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider’s entire set of resources. It MUST be a stable, non-reassignable identifier that does not change when the same resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. REQUIRED.",
						Schema:      "urn:scim:schemas:core:1.0",
						ReadOnly:    true,
						Required:    true,
					}, {
						Name:        "displayName",
						Type:        "string",
						Description: "A human readable name for the Group. REQUIRED.",
						Schema:      "urn:scim:schemas:core:1.0",
						ReadOnly:    true,
						Required:    true,
					}, {
						Name:                          "members",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "member",
						Description:                   "A list of members of the Group. Canonical Type \"User\" is currently supported and READ-ONLY. The value must be the \"id\" of a SCIM User resource.",
						Schema:                        "urn:scim:schemas:core:1.0",
						Required:                      true,
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "The \"id\" of a SCIM User resource. REQUIRED.",
								ReadOnly:    true,
								Required:    true,
							},
							{
								Name:        "display",
								Type:        "string",
								Description: "A human readable name for the member, primarily used for display purposes. READ-ONLY.",
								ReadOnly:    true,
							},
						},
					},
				},
			},
			body: `{
  "name": "Group",
  "description": "Core Group",
  "schema": "urn:scim:schemas:core:1.0",
  "endpoint": "/Groups",
  "attributes": [
    {
      "name": "id",
      "type": "string",
      "multiValued": false,
      "description": "Unique identifier for the SCIM resource as defined by the Service Provider. Each representation of the resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider’s entire set of resources. It MUST be a stable, non-reassignable identifier that does not change when the same resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. REQUIRED.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": true,
      "required": true,
      "caseExact": false
    },
    {
      "name": "displayName",
      "type": "string",
      "multiValued": false,
      "description": "A human readable name for the Group. REQUIRED.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": true,
      "required": true,
      "caseExact": false
    },
    {
      "name": "members",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "member",
      "description": "A list of members of the Group. Canonical Type \"User\" is currently supported and READ-ONLY. The value must be the \"id\" of a SCIM User resource.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": true,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "The \"id\" of a SCIM User resource. REQUIRED.",
          "readOnly": true,
          "required": true,
          "caseExact": false
        },
        {
          "name": "display",
          "type": "string",
          "multiValued": false,
          "description": "A human readable name for the member, primarily used for display purposes. READ-ONLY.",
          "readOnly": true,
          "required": false,
          "caseExact": false
        }
      ]
    }
  ]
}`,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get("/scim/v1/Schemas/Groups").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		schema, resp, err := client.GetGroupSchema(ctx)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, schema)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, schema)
		}
	}
}
