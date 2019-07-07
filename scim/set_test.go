package scim

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_SetHTTPClient(t *testing.T) {
	c := &Client{}

	c.SetHTTPClient(nil)
	require.Equal(t, http.DefaultClient, c.httpClient)

	c.SetHTTPClient(&http.Client{
		Timeout: 10 * time.Second,
	})
	require.Equal(t, &http.Client{
		Timeout: 10 * time.Second,
	}, c.httpClient)
}

func TestClient_SetParseResp(t *testing.T) {
	c := &Client{}

	c.SetParseResp(nil)
	require.NotNil(t, c.parseResp)

	c.SetParseResp(ParseRespDefault)
	require.NotNil(t, c.parseResp)
}

func TestClient_SetParseErrorResp(t *testing.T) {
	c := &Client{}

	c.SetParseErrorResp(nil)
	require.NotNil(t, c.parseErrorResp)

	c.SetParseErrorResp(ParseErrorRespDefault)
	require.NotNil(t, c.parseErrorResp)
}

func TestClient_SetIsError(t *testing.T) {
	c := &Client{}

	c.SetIsError(nil)
	require.NotNil(t, c.isError)

	c.SetIsError(IsErrorDefault)
	require.NotNil(t, c.isError)
}

func TestClient_SetEndpoint(t *testing.T) {
	c := &Client{}

	c.SetEndpoint("")
	require.Equal(t, DefaultEndpoint, c.endpoint)

	ep := "https://example.com"
	c.SetEndpoint(ep)
	require.Equal(t, ep, c.endpoint)
}
