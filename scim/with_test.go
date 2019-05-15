package scim

import (
	"testing"

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

func TestClientWithNewHTTPClient(t *testing.T) {
	c := &Client{}

	c2 := c.WithNewHTTPClient(nil)
	require.Nil(t, c.newHTTPClient)
	require.NotNil(t, c2.newHTTPClient)
}

func TestClientWithParseResp(t *testing.T) {
	c := &Client{}

	c2 := c.WithParseResp(nil)
	require.Nil(t, c.parseResp)
	require.NotNil(t, c2.parseResp)
}

func TestClientWithParseErrorResp(t *testing.T) {
	c := &Client{}

	c2 := c.WithParseErrorResp(nil)
	require.Nil(t, c.parseErrorResp)
	require.NotNil(t, c2.parseErrorResp)
}

func TestClientWithIsError(t *testing.T) {
	c := &Client{}

	c2 := c.WithIsError(nil)
	require.Nil(t, c.isError)
	require.NotNil(t, c2.isError)
}

func TestClientWithEndpoint(t *testing.T) {
	c := &Client{}

	ep := "endpoint"
	c2 := c.WithEndpoint(ep)
	require.Equal(t, "", c.endpoint)
	require.NotNil(t, ep, c2.endpoint)
}
