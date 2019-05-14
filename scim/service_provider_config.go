package scim

import (
	"context"
	"net/http"
)

type (
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

	AuthenticationScheme struct {
		Primary     bool   `json:"primary"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		SpecURL     string `json:"specUrl"`
	}

	BulkConfig struct {
		Supported       bool `json:"supported"`
		MaxOperations   int  `json:"maxOperations"`
		MaxPlayloadSize int  `json:"maxPlayloadSize"`
	}

	FilterConfig struct {
		Supported  bool `json:"supported"`
		MaxResults int  `json:"maxResults"`
	}

	PatchConfig struct {
		Supported bool `json:"supported"`
	}

	ChangePasswordConfig struct {
		Supported bool `json:"supported"`
	}

	SortConfig struct {
		Supported bool `json:"supported"`
	}

	EtagConfig struct {
		Supported bool `json:"supported"`
	}

	XMLDataFormatConfig struct {
		Supported bool `json:"supported"`
	}
)

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
func (c *client) GetServiceProviderConfigResp(ctx context.Context) (*http.Response, error) {
	// GET /ServiceProviderConfigs
	return c.getResp(ctx, "GET", "/ServiceProviderConfigs", nil, nil)
}
