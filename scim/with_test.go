package scim

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_copy(t *testing.T) {
	c := &Client{
		endpoint: "endpoint",
		token:    "token",
	}

	c2 := c.copy()
	require.Equal(t, c, c2)
	c2.endpoint = "change"
	c2.token = "change"
	require.NotEqual(t, c2.endpoint, c.endpoint)
	require.NotEqual(t, c2.token, c.token)
}

func TestClient_WithHTTPClient(t *testing.T) {
	c := &Client{}

	c2 := c.WithHTTPClient(nil)
	require.Nil(t, c.httpClient)
	require.Equal(t, http.DefaultClient, c2.httpClient)
	c3 := c.WithHTTPClient(&http.Client{
		Timeout: 10 * time.Second,
	})
	require.Nil(t, c.httpClient)
	require.Equal(t, &http.Client{
		Timeout: 10 * time.Second,
	}, c3.httpClient)
}

func TestClient_WithParseResp(t *testing.T) {
	c := &Client{}

	c2 := c.WithParseResp(nil)
	require.Nil(t, c.parseResp)
	require.NotNil(t, c2.parseResp)
}

func TestClient_WithParseErrorResp(t *testing.T) {
	c := &Client{}

	c2 := c.WithParseErrorResp(nil)
	require.Nil(t, c.parseErrorResp)
	require.NotNil(t, c2.parseErrorResp)
}

func TestClient_WithIsError(t *testing.T) {
	c := &Client{}

	c2 := c.WithIsError(nil)
	require.Nil(t, c.isError)
	require.NotNil(t, c2.isError)
}

func TestClient_WithEndpoint(t *testing.T) {
	c := &Client{}

	ep := "endpoint"
	c2 := c.WithEndpoint(ep)
	require.Equal(t, "", c.endpoint)
	require.NotNil(t, ep, c2.endpoint)
}
