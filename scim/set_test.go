package scim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_clientSetNewHTTPClient(t *testing.T) {
	c := &client{}

	c.SetNewHTTPClient(nil)
	require.NotNil(t, c.newHTTPClient)

	c.SetNewHTTPClient(NewHTTPClientDefault)
	require.NotNil(t, c.newHTTPClient)
}

func Test_clientSetParseResp(t *testing.T) {
	c := &client{}

	c.SetParseResp(nil)
	require.NotNil(t, c.parseResp)

	c.SetParseResp(ParseRespDefault)
	require.NotNil(t, c.parseResp)
}

func Test_clientSetParseErrorResp(t *testing.T) {
	c := &client{}

	c.SetParseErrorResp(nil)
	require.NotNil(t, c.parseErrorResp)

	c.SetParseErrorResp(ParseErrorRespDefault)
	require.NotNil(t, c.parseErrorResp)
}

func Test_clientSetIsError(t *testing.T) {
	c := &client{}

	c.SetIsError(nil)
	require.NotNil(t, c.isError)

	c.SetIsError(IsErrorDefault)
	require.NotNil(t, c.isError)
}

func Test_clientSetEndpoint(t *testing.T) {
	c := &client{}

	c.SetEndpoint("")
	require.Equal(t, DefaultEndpoint, c.endpoint)

	ep := "https://example.com"
	c.SetEndpoint(ep)
	require.Equal(t, ep, c.endpoint)
}
