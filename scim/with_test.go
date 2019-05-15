package scim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_client_copy(t *testing.T) {
	c := &client{
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

func Test_clientWithNewHTTPClient(t *testing.T) {
	c := &client{}

	c2 := c.WithNewHTTPClient(nil)
	c3 := c2.(*client)
	require.Nil(t, c.newHTTPClient)
	require.NotNil(t, c3.newHTTPClient)
}

func Test_clientWithParseResp(t *testing.T) {
	c := &client{}

	c2 := c.WithParseResp(nil)
	c3 := c2.(*client)
	require.Nil(t, c.parseResp)
	require.NotNil(t, c3.parseResp)
}

func Test_clientWithParseErrorResp(t *testing.T) {
	c := &client{}

	c2 := c.WithParseErrorResp(nil)
	c3 := c2.(*client)
	require.Nil(t, c.parseErrorResp)
	require.NotNil(t, c3.parseErrorResp)
}

func Test_clientWithIsError(t *testing.T) {
	c := &client{}

	c2 := c.WithIsError(nil)
	c3 := c2.(*client)
	require.Nil(t, c.isError)
	require.NotNil(t, c3.isError)
}

func Test_clientWithEndpoint(t *testing.T) {
	c := &client{}

	ep := "endpoint"
	c2 := c.WithEndpoint(ep)
	c3 := c2.(*client)
	require.Equal(t, "", c.endpoint)
	require.NotNil(t, ep, c3.endpoint)
}
