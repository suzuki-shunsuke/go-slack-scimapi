package scim

import (
	"context"
	"net/http"
)

type (
	// ServiceProviderConfig is Slack's configuration details for our SCIM API, including which operations are supported.
	// https://api.slack.com/scim#service_provider_configuration
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	ServiceProviderConfig struct {
		AuthenticationSchemes []AuthenticationScheme `json:"authenticationSchemes"`
		Patch                 *PatchConfig           `json:"patch"`
		Bulk                  *BulkConfig            `json:"bulk"`
		Filter                *FilterConfig          `json:"filter"`
		ChangePassword        *ChangePasswordConfig  `json:"changePassword"`
		Sort                  *SortConfig            `json:"sort"`
		Etag                  *EtagConfig            `json:"etag"`
		XMLDataFormat         *XMLDataFormatConfig   `json:"xmlDataFormat"`
	}

	// AuthenticationScheme specifies supported Authentication Scheme properties.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	AuthenticationScheme struct {
		Primary     bool   `json:"primary"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		SpecURL     string `json:"specUrl"`
	}

	// BulkConfig specifies BULK configuration options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	BulkConfig struct {
		Supported       bool `json:"supported"`
		MaxOperations   int  `json:"maxOperations"`
		MaxPlayloadSize int  `json:"maxPlayloadSize"`
	}

	// FilterConfig specifies FILTER options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	FilterConfig struct {
		Supported  bool `json:"supported"`
		MaxResults int  `json:"maxResults"`
	}

	// PatchConfig specifies PATCH configuration options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	PatchConfig struct {
		Supported bool `json:"supported"`
	}

	// ChangePasswordConfig specifies Change Password configuration options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	ChangePasswordConfig struct {
		Supported bool `json:"supported"`
	}

	// SortConfig specifies Sort configuration options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	SortConfig struct {
		Supported bool `json:"supported"`
	}

	// EtagConfig specifies Etag configuration options.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	EtagConfig struct {
		Supported bool `json:"supported"`
	}

	// XMLDataFormatConfig specifies whether the XML data format is supported.
	// http://www.simplecloud.info/specs/draft-scim-core-schema-01.html#rfc.section.9
	XMLDataFormatConfig struct {
		Supported bool `json:"supported"`
	}
)

// GetServiceProviderConfig sends GET ServiceProviderConfig API and returns ServiceProviderConfig .
// The returned response body is closed.
func (c *client) GetServiceProviderConfig(ctx context.Context) (*ServiceProviderConfig, *http.Response, error) {
	// GET /ServiceProviderConfigs
	resp, err := c.GetServiceProviderConfigResp(ctx)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	cfg := &ServiceProviderConfig{}
	return cfg, resp, c.parseResponse(resp, cfg)
}

// GetServiceProviderConfigResp sends GET service provider configuration API and returns an HTTP response.
// Internally, this method returns the returned values of *http.Client.Do .
func (c *client) GetServiceProviderConfigResp(ctx context.Context) (*http.Response, error) {
	// GET /ServiceProviderConfigs
	return c.getResp(ctx, "GET", "/ServiceProviderConfigs", nil, nil)
}
